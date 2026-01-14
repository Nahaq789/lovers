package userid

import "context"

type contextUserId struct{}

func WithContext(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, contextUserId{}, userId)
}

func FromContext(ctx context.Context) string {
	id, ok := ctx.Value(contextUserId{}).(string)
	if !ok {
		return ""
	}
	return id
}
