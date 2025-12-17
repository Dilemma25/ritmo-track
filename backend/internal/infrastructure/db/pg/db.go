package pg

import (
	"context"
	"fmt"
	"ritmotrack-backend/internal/infrastructure/config"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MainDB struct {
	pgxPool *pgxpool.Pool
	sq      sq.StatementBuilderType
}

func (ths *MainDB) GetPool() *pgxpool.Pool {
	return ths.pgxPool
}

func (ths *MainDB) GetSq() sq.StatementBuilderType { return ths.sq }

func NewMainDB(cfg *config.Config) (*MainDB, error) {
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort,
	))

	if err != nil {
		return nil, fmt.Errorf("failed to connect to PG: %w", err)
	}

	return &MainDB{
		pgxPool: pool,
		sq:      sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}
