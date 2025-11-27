package main

import (
	"context"
	"lovers/cmd/di/aws"
	"lovers/cmd/initialize"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := initialize.InitLogger()

	ctx := context.Background()

	aws, err := aws.Initialize(ctx, logger)
	if err != nil {
		logger.ErrorContext(ctx, "failed to init aws client", "error", err)
		return
	}

	authSet, err := initialize.InitAuth(ctx, logger, aws.Cognito)
	if err != nil {
		logger.ErrorContext(ctx, "failed to init auth", "error", err)
		return
	}

	r := gin.Default()
	Router(r, *authSet)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	serverCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	c := make(chan error, 1)
	go func() {
		c <- server.ListenAndServe()
	}()

	select {
	case err := <-c:
		if err != nil && err != http.ErrServerClosed {
			logger.ErrorContext(ctx, "HTTP server ListenAndServe", "error", err)
		}
	case <-serverCtx.Done():
		logger.Info("Server stopping")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.ErrorContext(shutdownCtx, "HTTP Server Shutdown", "error", err)
		}
	}
}
