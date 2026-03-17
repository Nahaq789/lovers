package logger

import (
	"context"
	"log/slog"
	"lovers/internal/shared/infrastructure/tracing/key"
	"os"
)

type ContextHandler struct {
	handler slog.Handler
}

func NewContextHandler(handler slog.Handler) *ContextHandler {
	return &ContextHandler{handler: handler}
}

func (h *ContextHandler) Handle(ctx context.Context, record slog.Record) error {
	var attrs []slog.Attr

	trace := key.NewContextTrace()
	if id := trace.GetValueFromCtx(ctx); id != "" {
		attrs = append(attrs, slog.String(string(trace.GetKey()), id))
	}

	if len(attrs) > 0 {
		record.AddAttrs(attrs...)
	}
	return h.handler.Handle(ctx, record)
}

func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewContextHandler(h.handler.WithAttrs(attrs))
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return NewContextHandler(h.handler.WithGroup(name))
}

type contextKey struct{}

func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}

func FromContext(ctx context.Context) *slog.Logger {
	l, ok := ctx.Value(contextKey{}).(*slog.Logger)
	if !ok {
		l = InitLogger()
	}

	return l
}

func InitLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	contextHandler := NewContextHandler(handler)
	logger := slog.New(contextHandler)
	slog.SetDefault(logger)

	return logger
}
