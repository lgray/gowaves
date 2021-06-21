package metamask

import (
	"context"
	"github.com/pkg/errors"
	"github.com/semrush/zenrpc/v2"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

func RunMetaMaskService(ctx context.Context, address string) error {
	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	rpc.Register("", MetaMask{}) // public
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	server := &http.Server{Addr: address, Handler: nil}

	go func() {
		zap.S().Info("shutting down metamask service...")
		err := server.Shutdown(ctx)
		if err != nil && !errors.Is(err, context.Canceled) {
			zap.S().Errorf("failed to shutdown metamask service: %v", err)
		}
	}()
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}