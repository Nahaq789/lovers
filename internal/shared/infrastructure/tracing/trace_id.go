package tracing

import "context"

type TracingID interface {
	GetKey() string
	GetValueFromCtx(ctx context.Context) string
	GenerateID() string
}
