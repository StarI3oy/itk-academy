package main

import (
	"context"
	"os/signal"
	"syscall"
	"wallet-service/internal/app"
	"wallet-service/internal/config"
	"wallet-service/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log := logger.NewLogger(zapcore.InfoLevel, true)

	defer log.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error("error with init config", zap.Error(err))
	}

	errCh := make(chan error, 5)

	serverApp, err := app.New(cfg, log)
	if err != nil {
		log.Fatal("error with init api gateway app", zap.Error(err))
	}
	go func() {
		errCh <- serverApp.Start()
	}()

	select {
	case <-ctx.Done():
		log.Info("shutdown signal received")
		serverApp.Stop()
	case err := <-errCh:
		log.Fatal("error with auth server app: ", zap.Error(err))
	}
}
