package boostrap

import (
	"log"
	"ritmotrack-backend/internal/application/usecase/auth"
	migrationUseCase "ritmotrack-backend/internal/application/usecase/migration"
	"ritmotrack-backend/internal/application/usecase/user"
	"ritmotrack-backend/internal/infrastructure/adapter_di"
	"ritmotrack-backend/internal/infrastructure/config"
	"ritmotrack-backend/internal/infrastructure/db/pg"
	"ritmotrack-backend/internal/infrastructure/db/repository"
	"ritmotrack-backend/internal/infrastructure/provider"
	"ritmotrack-backend/internal/presentation/cli"
	"ritmotrack-backend/internal/presentation/http/controller"
	"ritmotrack-backend/internal/presentation/http/middleware"
	"ritmotrack-backend/internal/presentation/http/responder"
	"ritmotrack-backend/internal/presentation/http/validator"
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
	injectMiddlewares(container)
	injectHttpControllers(container)
	injectCliControllers(container)

	log.Println("DI container created")

	return container
}

// infrastructure
func injectConfigs(container *adapter_di.Container) {
	container.Provide(config.NewConfig)
}

func injectDB(container *adapter_di.Container) {
	container.Provide(pg.NewMainDB)
}

func injectProviders(container *adapter_di.Container) {
	container.Provide(provider.NewMigrationProvider)

	container.Provide(provider.NewHasherProvider)

	container.Provide(provider.NewJwtProvider)
}

func injectRepositories(container *adapter_di.Container) {
	container.Provide(repository.NewUserRepository)
}

// application
func injectUseCases(container *adapter_di.Container) {
	//User
	container.Provide(user.NewUserCreateUseCase)
	container.Provide(user.NewUserGetByIDUseCase)
	container.Provide(user.NewUserGetAllUseCase)
	container.Provide(user.NewPatchUserUseCase)

	//Migration
	container.Provide(migrationUseCase.NewUpgradeMigrationUseCase)
	container.Provide(migrationUseCase.NewDownMigrationUseCase)
	container.Provide(migrationUseCase.NewResetMigrationUseCase)

	//Auth
	container.Provide(auth.NewCreateJwtUseCase)
	container.Provide(auth.NewCheckJwtUseCase)
}

// presentation
func injectMiddlewares(container *adapter_di.Container) {
	container.Provide(middleware.NewAuthMiddleware)
}

func injectHttpControllers(container *adapter_di.Container) {
	container.Provide(controller.NewUserController)

	container.Provide(controller.NewAuthController)

	container.Provide(responder.NewResponder)
	container.Provide(validator.NewValidator)
}

func injectCliControllers(container *adapter_di.Container) {
	container.Provide(cli.NewMigrationCommandController)
}
