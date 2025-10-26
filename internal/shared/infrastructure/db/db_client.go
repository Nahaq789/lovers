package db

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

type DbClient struct {
	logger *slog.Logger
	client *sql.DB
}

func InitDbClient(ctx context.Context, l *slog.Logger, connStr string) (*DbClient, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		l.ErrorContext(ctx, "failed to init db client", "error", err)
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(1 * time.Minute)

	// 接続確認
	if err := db.PingContext(ctx); err != nil {
		l.ErrorContext(ctx, "failed to ping db", "error", err)
		return nil, err
	}

	l.InfoContext(ctx, "db client initialized successfully")

	return &DbClient{
		logger: l,
		client: db,
	}, nil
}

func (d *DbClient) Close(ctx context.Context) error {
	return d.client.Close()
}
