package di

import (
	"ritmotrack-backend/internal/application/user/usecase"
	"ritmotrack-backend/internal/core/config"
	repoInterface "ritmotrack-backend/internal/domain/repository"
	repoImpl "ritmotrack-backend/internal/infrastructure/db/repository"
	"ritmotrack-backend/internal/presentation/http/controller"
)

func CreateContainer() *Container {
	container := NewContainer()

	//infrastructure
	injectConfigs(container)
	injectRepositories(container)

	//application
	injectUseCases(container)

	//presentation
	injectControllers(container)

	return container
}

// infrastructure
func injectConfigs(container *Container) {
	container.Provide(config.NewConfig)
}

func injectRepositories(container *Container) {
	container.Provide(func() repoInterface.UserRepository {
		return repoImpl.UserRepository{}
	})
}

// application
func injectUseCases(container *Container) {
	container.Provide(usecase.NewUserCreateUseCase)
	container.Provide(usecase.NewUserGetByIDUseCase)
	container.Provide(usecase.NewUserGetAllUseCase)
}

// presentation
func injectControllers(container *Container) {
	container.Provide(controller.NewUserController)
}
