package main

import (
	"context"
	"errors"
	"log/slog"
	"lovers/internal/shared/infrastructure/logger"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	contextHandler := logger.NewContextHandler(handler)
	logger := slog.New(contextHandler)
	slog.SetDefault(logger)

	r := gin.Default()
	Router(r)
	server := &http.Server{Addr: ":8080", Handler: r.Handler()}

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.InfoContext(ctx, "HTTP Server Shutdown", "error", err)
		}
		wg.Done()
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.ErrorContext(context.Background(), "HTTP server ListenAndServe", "error", err)
	}
}
