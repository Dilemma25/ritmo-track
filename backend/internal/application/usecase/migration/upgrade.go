package migration

import "ritmotrack-backend/internal/application/provider"

type UpgradeMigrationUseCase interface {
	Execute() error
}

type upgradeMigrationUseCase struct {
	migrationProvider provider.MigrationProvider
}

func NewUpgradeMigrationUseCase(migrationProvider provider.MigrationProvider) UpgradeMigrationUseCase {
	return &upgradeMigrationUseCase{migrationProvider: migrationProvider}
}

func (ths *upgradeMigrationUseCase) Execute() error {
	return ths.migrationProvider.Upgrade()
}
