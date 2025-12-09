package di

import (
	"ritmotrack-backend/internal/application/user/usecase"
	"ritmotrack-backend/internal/core/config"
	repoInterface "ritmotrack-backend/internal/domain/repository"
	repoImpl "ritmotrack-backend/internal/infrastructure/db/repository"
	httpLayer "ritmotrack-backend/internal/presentation/http/controller"
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
	container.Provide(func(r repoInterface.UserRepository) usecase.UserCreateUseCase {
		return usecase.UserCreateUseCase{
			Repo: r,
		}
	})

	container.Provide(func(r repoInterface.UserRepository) usecase.UserGetByIDUseCase {
		return usecase.UserGetByIDUseCase{
			Repo: r,
		}
	})

	container.Provide(func(r repoInterface.UserRepository) usecase.UserGetAllUseCase {
		return usecase.UserGetAllUseCase{
			Repo: r,
		}
	})
}

// presentation
func injectControllers(container *Container) {
	container.Provide(func(
		userCreateUseCase usecase.UserCreateUseCase,
		userGetByIdUseCase usecase.UserGetByIDUseCase,
		userGetAllUseCase usecase.UserGetAllUseCase,
	) *httpLayer.UserController {
		return &httpLayer.UserController{
			UserCreateUseCase:  userCreateUseCase,
			UserGetByIDUseCase: userGetByIdUseCase,
			UserGetAllUseCase:  userGetAllUseCase,
		}
	})
}
