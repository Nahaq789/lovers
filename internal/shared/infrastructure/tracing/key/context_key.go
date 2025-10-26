package key

import (
	"context"

	"github.com/google/uuid"
)

type ContextKey string

const (
	ContextTraceKey ContextKey = "traceID"
)

type ContextTrace struct {
	key ContextKey
}

func NewContextTrace() ContextTrace {
	return ContextTrace{key: ContextTraceKey}
}

func (c ContextTrace) GetKey() ContextKey {
	return c.key
}

func (c ContextTrace) GetValueFromCtx(ctx context.Context) string {
	if id, ok := ctx.Value(c.GetKey()).(string); ok && id != "" {
		return id
	}
	return ""
}

func (c ContextTrace) GenerateID() string {
	id := uuid.New().String()
	return id
}
