package main

import (
	"log/slog"
	"lovers/internal/shared/infrastructure/logger"
	"net/http"
	"os"

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
	http.ListenAndServe(":8080", r)
}
