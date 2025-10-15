package repositories

import "context"

type UserRepository interface {
	Register(ctx context.Context)
}
