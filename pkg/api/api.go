package api

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ic-n/sx/contracts"
	commitrevealv1 "github.com/ic-n/sx/pkg/api/gen/commitreveal/v1"
	"github.com/ic-n/sx/pkg/tools"
)

type Server struct {
	log          *slog.Logger
	auth         *bind.TransactOpts // todo add context
	client       bind.ContractBackend
	phaseSeconds int
	commitrevealv1.UnimplementedCommitRevealServiceServer
}

func NewServer(
	log *slog.Logger,
	auth *bind.TransactOpts,
	client bind.ContractBackend,
	phaseSeconds int,
) *Server {
	return &Server{
		log:          log,
		auth:         auth,
		client:       client,
		phaseSeconds: phaseSeconds,
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
	if _, err := s.client.SuggestGasPrice(ctx); err != nil {
		return &commitrevealv1.HealthResponse{
			Ok: false,
		}, fmt.Errorf("blockchain is not online: %w", err)
	}

	return &commitrevealv1.HealthResponse{
		Ok: true,
	}, nil
}

func (s Server) CreatePoll(_ context.Context, r *commitrevealv1.CreatePollRequest) (*commitrevealv1.CreatePollResponse, error) {
	addr, _, _, err := contracts.DeployCommitReveal(s.auth, s.client, big.NewInt(int64(s.phaseSeconds)), r.GetChoice_1(), r.GetChoice_2())
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contract: %w", err)
	}

	return &commitrevealv1.CreatePollResponse{
		Address: addr.Hex(),
	}, nil
}

func (s Server) GetPoll(_ context.Context, r *commitrevealv1.GetPollRequest) (*commitrevealv1.GetPollResponse, error) {
	c, err := s.contract(r.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("failed to find initilise contract bindings: %w", err)
	}

	commitEnd, err := c.CommitPhaseEndTime(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get phase end time: %w", err)
	}

	commitEndTime := time.Unix(commitEnd.Int64(), 0)

	choice1, err := c.Choice1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get choice1: %w", err)
	}

	choice2, err := c.Choice2(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get choice2: %w", err)
	}

	votesForChoice1, err := c.VotesForChoice1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get choice1 count: %w", err)
	}

	votesForChoice2, err := c.VotesForChoice2(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get choice2 count: %w", err)
	}

	return &commitrevealv1.GetPollResponse{
		SecondsLeft: int64(time.Until(commitEndTime).Seconds()),
		Choice_1:    choice1,
		Choice_2:    choice2,
		Count_1:     votesForChoice1.Int64(),
		Count_2:     votesForChoice2.Int64(),
	}, nil
}

func (s Server) Commit(_ context.Context, r *commitrevealv1.CommitRequest) (*commitrevealv1.CommitResponse, error) {
	c, err := s.contract(r.GetAddress())
	if err != nil {
		return &commitrevealv1.CommitResponse{
			Ok: false,
		}, fmt.Errorf("failed to find initilise contract bindings: %w", err)
	}

	if _, err := c.CommitVote(s.auth, tools.B32(r.SecretHash)); err != nil {
		return &commitrevealv1.CommitResponse{
			Ok: false,
		}, fmt.Errorf("failed to commit vote: %w", err)
	}

	return &commitrevealv1.CommitResponse{
		Ok: true,
	}, nil
}

func (s Server) Reveal(_ context.Context, r *commitrevealv1.RevealRequest) (*commitrevealv1.RevealResponse, error) {
	c, err := s.contract(r.GetAddress())
	if err != nil {
		return &commitrevealv1.RevealResponse{
			Ok: false,
		}, fmt.Errorf("failed to find initilise contract bindings: %w", err)
	}

	if _, err := c.RevealVote(s.auth, r.Secret, tools.B32(r.SecretHash)); err != nil {
		return &commitrevealv1.RevealResponse{
			Ok: false,
		}, fmt.Errorf("failed to commit vote: %w", err)
	}

	return &commitrevealv1.RevealResponse{
		Ok: true,
	}, nil
}

func (s Server) contract(addr string) (*contracts.CommitReveal, error) {
	return contracts.NewCommitReveal(common.HexToAddress(addr), s.client)
}
