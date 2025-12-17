package cli

import (
	"fmt"
	"ritmotrack-backend/internal/application/usecase/migration"
)

type MigrationCommandController interface {
	Upgrade()
}

type migrationCommandController struct {
	upgradeMigrationUseCase migration.UpgradeMigrationUseCase
}

func NewMigrationCommandController(
	upgradeMigrationUseCase migration.UpgradeMigrationUseCase,
) MigrationCommandController {
	return migrationCommandController{
		upgradeMigrationUseCase: upgradeMigrationUseCase,
	}
}

func (ths migrationCommandController) Upgrade() {
	if err := ths.upgradeMigrationUseCase.Execute(); err != nil {
		fmt.Println(err)
	}
}
