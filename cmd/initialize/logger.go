package initialize

import (
	"context"
	"log/slog"
	"lovers/internal/shared/infrastructure/logger"
)

func InitLogger() *slog.Logger {
	return logger.InitLogger()
}

func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return logger.WithContext(ctx, l)
}
