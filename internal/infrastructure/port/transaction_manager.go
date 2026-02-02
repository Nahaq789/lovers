package port

import (
	"context"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/transaction"
)

type TransactionManagerImpl struct {
	db *db.DbClient
}

func NewTransactionManager(d *db.DbClient) *TransactionManagerImpl {
	return &TransactionManagerImpl{db: d}
}

func (tm *TransactionManagerImpl) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.db.GetClient().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	txCtx := transaction.WithContext(ctx, tx)

	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
