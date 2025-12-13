package cli

import (
	"fmt"
	"ritmotrack-backend/internal/application/migration/usecase"
)

type MigrationCommandController interface {
	Upgrade()
}

type migrationCommandController struct {
	upgradeMigrationUseCase usecase.UpgradeMigrationUseCase
}

func NewMigrationCommandController(
	upgradeMigrationUseCase usecase.UpgradeMigrationUseCase,
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
