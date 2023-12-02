package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ic-n/sx/pkg/api"
	"github.com/ic-n/sx/pkg/tools"
	"github.com/urfave/cli/v2"
)

var log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

var (
	netURL = &cli.StringFlag{
		Name: "net-url",
	}
	keyPath = &cli.StringFlag{
		Name: "private-key",
	}
	chainID = &cli.Int64Flag{
		Name: "chain-id",
	}
)

func main() {
	app := &cli.App{
		Name: "commit-reveal-server",
		Flags: []cli.Flag{
			netURL, keyPath, chainID,
		},
		Action: App,
	}

	if err := app.Run(os.Args); err != nil {
		log.Error("failed to run app", "error", err)
	}
}

func App(cctx *cli.Context) error {
	ctx, cancel := context.WithTimeout(cctx.Context, 10*time.Minute)
	defer cancel()

	client, err := ethclient.DialContext(ctx, netURL.Get(cctx))
	if err != nil {
		return fmt.Errorf("failed to dial eth network: %w", err)
	}

	pk, err := tools.ReadPrivateKeyFile(keyPath.Get(cctx))
	if err != nil {
		return fmt.Errorf("failed to read private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(chainID.Get(cctx)))
	if err != nil {
		return fmt.Errorf("failed to initiate transactor: %w", err)
	}

	h, err := api.NewHandler(ctx, api.NewServer(log, auth.From, auth.Signer, client))
	if err != nil {
		return fmt.Errorf("failed to create server handler: %w", err)
	}

	s := http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}),
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 5,
	}
	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
