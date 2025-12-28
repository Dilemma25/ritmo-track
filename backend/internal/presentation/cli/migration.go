package cli

import (
	"fmt"
	"ritmotrack-backend/internal/application/usecase/migration"
)

type MigrationCommandController interface {
	Upgrade()
	Down()
	Reset()
}

type migrationCommandController struct {
	upgradeMigrationUseCase migration.UpgradeMigrationUseCase
	downMigrationUseCase    migration.DownMigrationUseCase
	resetMigrationUseCase   migration.ResetMigrationUseCase
}

func NewMigrationCommandController(
	upgradeMigrationUseCase migration.UpgradeMigrationUseCase,
	downMigrationUseCase migration.DownMigrationUseCase,
	resetMigrationUseCase migration.ResetMigrationUseCase,
) MigrationCommandController {
	return &migrationCommandController{
		upgradeMigrationUseCase: upgradeMigrationUseCase,
		downMigrationUseCase:    downMigrationUseCase,
		resetMigrationUseCase:   resetMigrationUseCase,
	}
}

func (ths *migrationCommandController) Upgrade() {
	if err := ths.upgradeMigrationUseCase.Execute(); err != nil {
		fmt.Println(err)
	}
}

func (ths *migrationCommandController) Down() {
	if err := ths.downMigrationUseCase.Execute(); err != nil {
		fmt.Println(err)
	}
}

func (ths *migrationCommandController) Reset() {
	fmt.Println("\033[31mAre you sure you want to reset all migrations?\033[0m")
	fmt.Print("\033[31mType 'reset-all-migrations' to confirm: \033[0m")

	var input string

	_, _ = fmt.Scan(&input)

	if input != "reset-all-migrations" {
		fmt.Println("Migrations reset cancelled")
		return
	}

	if err := ths.resetMigrationUseCase.Execute(); err != nil {
		fmt.Println(err)
	}
}
