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

func (ths *GetUserByIDUseCase) Execute(ctx context.Context, id uint) (*dto.UserDTO, error) {
	user, _ := ths.repo.GetById(ctx, id)

	userDTO := &dto.UserDTO{
		Id:    user.GetId(),
		Name:  user.GetName(),
		Login: user.GetLogin(),
	}

	return userDTO, nil
}
