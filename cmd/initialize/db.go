package initialize

import (
	"context"
	"fmt"
	"log/slog"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/sharedAws"
	"os"
)

func InitDB(ctx context.Context, l *slog.Logger, parameterStore *sharedAws.ParameterStoreClient) (*db.DbClient, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := db.InitDbClient(ctx, l, connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
