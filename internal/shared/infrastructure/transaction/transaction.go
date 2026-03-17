package transaction

import (
	"context"
	"database/sql"
)

type contextTransactionKey struct{}

func WithContext(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, contextTransactionKey{}, tx)
}

func FromContext(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(contextTransactionKey{}).(*sql.Tx)
	if !ok {
		return nil
	}
	return tx
}
