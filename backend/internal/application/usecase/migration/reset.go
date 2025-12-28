package migration

import "ritmotrack-backend/internal/application/provider"

type ResetMigrationUseCase interface {
	Execute() error
}

type resetMigrationUseCase struct {
	migrationProvider provider.MigrationProvider
}

func NewResetMigrationUseCase(migrationProvider provider.MigrationProvider) ResetMigrationUseCase {
	return &resetMigrationUseCase{
		migrationProvider: migrationProvider,
	}
}

func (ths *resetMigrationUseCase) Execute() error {
	return ths.migrationProvider.Reset()
}
