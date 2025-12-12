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
	"github.com/joho/godotenv"
)

func main() {
	l := initialize.InitLogger()

	envErr := godotenv.Load()
	if envErr != nil {
		l.Error("Error loading env")
		return
	}

	ctx := context.Background()
	ctxWithLogger := initialize.WithContext(ctx, l)

	aws, err := aws.Initialize(ctxWithLogger)
	if err != nil {
		l.ErrorContext(ctxWithLogger, "failed to init aws client", "error", err)
		return
	}

	authSet, err := initialize.InitAuth(ctxWithLogger, aws.Cognito)
	if err != nil {
		l.ErrorContext(ctxWithLogger, "failed to init auth", "error", err)
		return
	}

	r := gin.Default()
	r.ContextWithFallback = true
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
			l.ErrorContext(ctx, "HTTP server ListenAndServe", "error", err)
		}
	case <-serverCtx.Done():
		l.Info("Server stopping")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			l.ErrorContext(shutdownCtx, "HTTP Server Shutdown", "error", err)
		}
	}
}
