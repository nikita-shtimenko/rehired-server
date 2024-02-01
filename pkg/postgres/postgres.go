package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(connectionString string, timeout time.Duration) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
