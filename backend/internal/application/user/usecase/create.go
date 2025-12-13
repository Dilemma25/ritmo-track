package usecase

import (
	"context"
	"ritmotrack-backend/internal/application/user/dto"
	"ritmotrack-backend/internal/domain/entity"
	"ritmotrack-backend/internal/domain/repository"
	"time"
)

type UserCreateUseCase struct {
	repo repository.UserRepository
}

func NewUserCreateUseCase(repo repository.UserRepository) *UserCreateUseCase {
	return &UserCreateUseCase{repo: repo}
}

func (ths UserCreateUseCase) Execute(ctx context.Context, data dto.CreateUserDTO) (dto.UserOutputDTO, error) {
	user := entity.User{
		Name:      data.Name,
		Login:     data.Login,
		Password:  data.Password,
		CreatedAt: time.Now(),
	}

	newUser, _ := ths.repo.Create(ctx, user)

	newUserDTO := dto.UserOutputDTO{
		Id:    newUser.Id,
		Name:  newUser.Name,
		Login: newUser.Login,
	}

	return newUserDTO, nil
}
