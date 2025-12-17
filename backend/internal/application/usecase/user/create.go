package user

import (
	"context"
	userDTO "ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/provider"
	"ritmotrack-backend/internal/domain/entity"
	"ritmotrack-backend/internal/domain/repository"
)

type CreateUserUseCase struct {
	repo   repository.UserRepository
	hasher provider.HasherProvider
}

func NewUserCreateUseCase(
	repo repository.UserRepository,
	hasher provider.HasherProvider,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (ths CreateUserUseCase) Execute(ctx context.Context, data userDTO.CreateUserDTO) (userDTO.UserOutputDTO, error) {

	user := entity.NewUser(data.Name, data.Login, ths.hasher.Hash(data.Password))

	if err := ths.repo.Create(ctx, user); err != nil {
		return userDTO.UserOutputDTO{}, err
	}

	newUserDTO := userDTO.UserOutputDTO{
		Id:        user.GetId(),
		Name:      user.GetName(),
		Login:     user.GetLogin(),
		CreatedAt: user.GetCreatedAt(),
	}

	return newUserDTO, nil
}
