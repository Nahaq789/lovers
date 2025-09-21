package logger

import (
	"context"
	"log/slog"
	"lovers/internal/shared/infrastructure/tracing/key"
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
		attrs = append(attrs, slog.String(trace.GetKey(), id))
	}

	if len(attrs) > 0 {
		record.AddAttrs(attrs...)
	}
	return h.handler.Handle(ctx, record)
}
