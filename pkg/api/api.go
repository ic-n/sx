package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	commitrevealv1 "github.com/ic-n/sx/pkg/api/gen/commitreveal/v1"
)

type Server struct {
	log *slog.Logger
	// auth *bind.TransactOpts
	// cli  blockchain.Client
	// c    *contracts.CommitReveal
	commitrevealv1.UnimplementedCommitRevealServiceServer
}

func (s Server) Health(_ context.Context, _ *commitrevealv1.HealthRequest) (*commitrevealv1.HealthResponse, error) {
	return &commitrevealv1.HealthResponse{
		Ok: true,
	}, nil
}

func NewServer(log *slog.Logger) *Server {
	return &Server{log: log}
}

func NewHandler(ctx context.Context, s *Server) (http.Handler, error) {
	mux := runtime.NewServeMux()
	if err := commitrevealv1.RegisterCommitRevealServiceHandlerServer(ctx, mux, s); err != nil {
		return nil, fmt.Errorf("failed to register service: %w", err)
	}

	return mux, nil
}
