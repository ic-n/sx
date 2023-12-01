package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ic-n/sx/pkg/api"
	"github.com/ic-n/sx/pkg/tools"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()

	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	errLog := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	client, from, signer, err := getSimulatedBackend()
	if err != nil {
		errLog.Error("failed to initialise simulated backend: %v", err)
		os.Exit(1)
	}

	h, err := api.NewHandler(ctx, api.NewServer(log, from, signer, client))
	if err != nil {
		errLog.Error("failed to create server handler: %s", err)
		os.Exit(1)
	}

	s := http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer client.Commit()
			h.ServeHTTP(w, r)
		}),
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 5,
	}
	if err := s.ListenAndServe(); err != nil {
		errLog.Error("failed to serve: %s", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func getSimulatedBackend() (*backends.SimulatedBackend, common.Address, bind.SignerFn, error) {
	const (
		simChainID = 1337
		gasLim     = 4712388
	)

	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, nil, fmt.Errorf("failed to generate key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(simChainID))
	if err != nil {
		return nil, common.Address{}, nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: tools.Wei(1000)}

	client := backends.NewSimulatedBackend(alloc, gasLim)

	return client, auth.From, auth.Signer, nil
}
