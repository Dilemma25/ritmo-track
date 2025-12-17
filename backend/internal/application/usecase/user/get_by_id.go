package user

import (
	"context"
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/domain/repository"
)

type GetUserByIDUseCase struct {
	repo repository.UserRepository
}

func NewUserGetByIDUseCase(repo repository.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{repo: repo}
}

func (ths *GetUserByIDUseCase) Execute(ctx context.Context, id int64) (dto.UserOutputDTO, error) {
	user, _ := ths.repo.GetById(ctx, id)

	userDTO := dto.UserOutputDTO{
		Id:    user.GetId(),
		Name:  user.GetName(),
		Login: user.GetLogin(),
	}

	return userDTO, nil
}
