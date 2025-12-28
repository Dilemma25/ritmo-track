package migration

import "ritmotrack-backend/internal/application/provider"

type DownMigrationUseCase interface {
	Execute() error
}

type downMigrationUseCase struct {
	migrationProvider provider.MigrationProvider
}

func NewDownMigrationUseCase(migrationProvider provider.MigrationProvider) DownMigrationUseCase {
	return &downMigrationUseCase{
		migrationProvider: migrationProvider,
	}
}

func (ths *downMigrationUseCase) Execute() error {
	return ths.migrationProvider.Down()
}
