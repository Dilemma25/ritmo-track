package usecase

import (
	"context"
	"ritmotrack-backend/internal/application/user/dto"
	"ritmotrack-backend/internal/domain/repository"
)

type UserGetByIDUseCase struct {
	repo repository.UserRepository
}

func NewUserGetByIDUseCase(repo repository.UserRepository) *UserGetByIDUseCase {
	return &UserGetByIDUseCase{repo: repo}
}

func (ths *UserGetByIDUseCase) Execute(ctx context.Context, id int64) (dto.UserOutputDTO, error) {
	user, _ := ths.repo.GetById(ctx, id)

	userDTO := dto.UserOutputDTO{
		Id:    user.Id,
		Name:  user.Name,
		Login: user.Login,
	}

	return userDTO, nil
}
