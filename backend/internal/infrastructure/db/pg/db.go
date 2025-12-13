package pg

import (
	"database/sql"
	"fmt"
	"ritmotrack-backend/internal/infrastructure/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	db *sql.DB
}

func (ths *DB) GetDB() *sql.DB {
	return ths.db
}

func NewPostgresDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName,
	)

	// используем драйвер "pgx"
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	// Настройки пула соединений
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	// Проверка соединения
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &DB{db: conn}, nil
}
