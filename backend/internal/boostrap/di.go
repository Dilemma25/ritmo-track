package boostrap

import (
	"log"
	migrationUseCase "ritmotrack-backend/internal/application/migration/usecase"
	userUseCase "ritmotrack-backend/internal/application/user/usecase"
	"ritmotrack-backend/internal/infrastructure/adapter_di"
	"ritmotrack-backend/internal/infrastructure/config"
	"ritmotrack-backend/internal/infrastructure/db/pg"
	"ritmotrack-backend/internal/infrastructure/db/repository"
	"ritmotrack-backend/internal/infrastructure/provider"
	"ritmotrack-backend/internal/presentation/cli"
	"ritmotrack-backend/internal/presentation/http/controller"
)

var container *adapter_di.Container

func GetContainer() *adapter_di.Container {
	if container == nil {
		container = CreateContainer()
	}

	return container
}

func CreateContainer() *adapter_di.Container {
	container = adapter_di.NewContainer()

	//infrastructure
	injectConfigs(container)
	injectDB(container)
	injectProviders(container)
	injectRepositories(container)

	//application
	injectUseCases(container)

	//presentation
	injectServerControllers(container)
	injectCliControllers(container)

	log.Println("DI container created")

	return container
}

// infrastructure
func injectConfigs(container *adapter_di.Container) {
	container.Provide(config.NewConfig)
}

func injectDB(container *adapter_di.Container) {
	container.Provide(pg.NewPostgresDB)
}

func injectProviders(container *adapter_di.Container) {
	container.Provide(provider.NewMigrationProvider)
}

func injectRepositories(container *adapter_di.Container) {
	container.Provide(repository.NewUserRepository)
}

// application
func injectUseCases(container *adapter_di.Container) {
	container.Provide(userUseCase.NewUserCreateUseCase)
	container.Provide(userUseCase.NewUserGetByIDUseCase)
	container.Provide(userUseCase.NewUserGetAllUseCase)

	container.Provide(migrationUseCase.NewUpgradeMigrationUseCase)
}

// presentation
func injectServerControllers(container *adapter_di.Container) {
	container.Provide(controller.NewUserController)
}

func injectCliControllers(container *adapter_di.Container) {
	container.Provide(cli.NewMigrationCommandController)
}
