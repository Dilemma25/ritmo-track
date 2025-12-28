package provider

import (
	"database/sql"
	"ritmotrack-backend/asset"
	"ritmotrack-backend/internal/application/provider"
	"ritmotrack-backend/internal/infrastructure/config"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/pressly/goose/v3"
)

type migrationProvider struct {
	config *config.Config
}

func NewMigrationProvider(config *config.Config) provider.MigrationProvider {
	goose.SetBaseFS(asset.MigrationFS)
	err := goose.SetDialect("postgres")
	if err != nil {
		panic(err)
	}

	return &migrationProvider{config: config}
}

func (ths *migrationProvider) getDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", ths.config.DbDNS())

	//db, err := sql.Open("pgx", fmt.Sprintf(
	//	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	ths.config.DbHost, ths.config.DbPort, ths.config.DbUser,
	//	ths.config.DbPassword, ths.config.DbName,
	//))
	//
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (ths *migrationProvider) Upgrade() error {
	db, err := ths.getDB()

	if err != nil {
		return err
	}

	if err = goose.Up(db, config.MigrationDir); err != nil {
		_ = db.Close()
		return err
	}

	return db.Close()
}

func (ths *migrationProvider) Down() error {
	db, err := ths.getDB()

	if err != nil {
		return err
	}

	if err = goose.Down(db, config.MigrationDir); err != nil {
		_ = db.Close()
		return err
	}

	return db.Close()
}

func (ths *migrationProvider) Reset() error {
	db, err := ths.getDB()

	if err != nil {
		return err
	}

	if err = goose.Reset(db, config.MigrationDir); err != nil {
		_ = db.Close()
		return err
	}

	return db.Close()
}
