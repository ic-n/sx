package api

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	commitrevealv1 "github.com/ic-n/sx/pkg/api/gen/commitreveal/v1"
	"github.com/ic-n/sx/pkg/contracts"
	"github.com/ic-n/sx/pkg/tools"
)

type Server struct {
	log    *slog.Logger
	auth   *bind.TransactOpts
	client bind.ContractBackend
	commitrevealv1.UnimplementedCommitRevealServiceServer
}

func NewServer(
	log *slog.Logger,
	auth *bind.TransactOpts,
	client bind.ContractBackend,
) *Server {
	return &Server{
		log:    log,
		auth:   auth,
		client: client,
	}
}

func NewHandler(ctx context.Context, s *Server) (http.Handler, error) {
	mux := runtime.NewServeMux()
	if err := commitrevealv1.RegisterCommitRevealServiceHandlerServer(ctx, mux, s); err != nil {
		return nil, fmt.Errorf("failed to register service: %w", err)
	}

	return mux, nil
}

func (s Server) Health(ctx context.Context, _ *commitrevealv1.HealthRequest) (*commitrevealv1.HealthResponse, error) {
	if _, err := s.client.HeaderByNumber(ctx, nil); err != nil {
		err = fmt.Errorf("blockchain is not online: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	return &commitrevealv1.HealthResponse{
		Ok: true,
	}, nil
}

func (s Server) CreatePoll(ctx context.Context, r *commitrevealv1.CreatePollRequest) (*commitrevealv1.CreatePollResponse, error) {
	addr, _, _, err := contracts.DeployCommitReveal(s.writeAuth(ctx), s.client, big.NewInt(r.GetSeconds()), r.GetChoice_1(), r.GetChoice_2())
	if err != nil {
		err = fmt.Errorf("failed to deploy contract: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	return &commitrevealv1.CreatePollResponse{
		Address: addr.Hex(),
	}, nil
}

func (s Server) GetPoll(ctx context.Context, r *commitrevealv1.GetPollRequest) (*commitrevealv1.GetPollResponse, error) {
	c, err := s.contract(r.GetAddress())
	if err != nil {
		err = fmt.Errorf("failed to find initilise contract bindings: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	h, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		err = fmt.Errorf("failed to get header: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	commitEnd, err := c.CommitPhaseEndTime(s.readAuth(ctx))
	if err != nil {
		err = fmt.Errorf("failed to get phase end time: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	choice1, err := c.Choice1(s.readAuth(ctx))
	if err != nil {
		err = fmt.Errorf("failed to get choice1: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	choice2, err := c.Choice2(s.readAuth(ctx))
	if err != nil {
		err = fmt.Errorf("failed to get choice2: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	votesForChoice1, err := c.VotesForChoice1(s.readAuth(ctx))
	if err != nil {
		err = fmt.Errorf("failed to get choice1 count: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	votesForChoice2, err := c.VotesForChoice2(s.readAuth(ctx))
	if err != nil {
		err = fmt.Errorf("failed to get choice2 count: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	return &commitrevealv1.GetPollResponse{
		SecondsLeft: max(0, commitEnd.Int64()-int64(h.Time)),
		Choice_1:    choice1,
		Choice_2:    choice2,
		Count_1:     votesForChoice1.Int64(),
		Count_2:     votesForChoice2.Int64(),
	}, nil
}

func (s Server) Commit(ctx context.Context, r *commitrevealv1.CommitRequest) (*commitrevealv1.CommitResponse, error) {
	c, err := s.contract(r.GetAddress())
	if err != nil {
		err = fmt.Errorf("failed to find initilise contract bindings: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	if _, err := c.CommitVote(s.writeAuth(ctx), tools.Keccak256(r.Secret)); err != nil {
		err = fmt.Errorf("failed to commit vote: %w", err)
		s.log.Warn(err.Error())
		return nil, &runtime.HTTPStatusError{
			HTTPStatus: http.StatusBadRequest,
			Err:        err,
		}
	}

	return &commitrevealv1.CommitResponse{
		Ok: true,
	}, nil
}

func (s Server) Reveal(ctx context.Context, r *commitrevealv1.RevealRequest) (*commitrevealv1.RevealResponse, error) {
	c, err := s.contract(r.GetAddress())
	if err != nil {
		err = fmt.Errorf("failed to find initilise contract bindings: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	if _, err := c.RevealVote(s.writeAuth(ctx), r.Secret, tools.Keccak256(r.Secret)); err != nil {
		err = fmt.Errorf("failed to reveal vote: %w", err)
		s.log.Warn(err.Error())
		return nil, &runtime.HTTPStatusError{
			HTTPStatus: http.StatusBadRequest,
			Err:        err,
		}
	}

	return &commitrevealv1.RevealResponse{
		Ok: true,
	}, nil
}

func (s Server) writeAuth(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:      s.auth.From,
		Nonce:     s.auth.Nonce,
		Signer:    s.auth.Signer,
		Value:     common.Big0,
		GasPrice:  s.auth.GasPrice,
		GasFeeCap: s.auth.GasFeeCap,
		GasTipCap: s.auth.GasTipCap,
		GasLimit:  s.auth.GasLimit,
		Context:   ctx,
		NoSend:    s.auth.NoSend,
	}
}

func (s Server) readAuth(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{
		Pending: true,
		From:    s.auth.From,
		Context: ctx,
	}
}

func (s Server) contract(addr string) (*contracts.CommitReveal, error) {
	return contracts.NewCommitReveal(common.HexToAddress(addr), s.client)
}
