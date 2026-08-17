package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wesleyfebarretos/argocd-applications/internal/config"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	dbConfig := config.Get().Database

	config, err := pgxpool.ParseConfig(dbConfig.ConnURI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database configuration: %w", err)
	}

	config.MaxConns = int32(dbConfig.MaxConns)

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	return pool, nil
}
