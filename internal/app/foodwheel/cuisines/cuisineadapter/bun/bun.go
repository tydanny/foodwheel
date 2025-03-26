package bun

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewBunDB(ctx context.Context, config Config) (*bun.DB, error) {
	poolConfig, err := pgxpool.ParseConfig(config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx pool config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx connection pool: %w", err)
	}

	return bun.NewDB(stdlib.OpenDBFromPool(pool), pgdialect.New()), nil
}
