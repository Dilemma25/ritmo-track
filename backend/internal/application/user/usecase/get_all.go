package usecase

import (
	"context"
	"ritmotrack-backend/internal/application/user/dto"
	"ritmotrack-backend/internal/domain/repository"
)

type UserGetAllUseCase struct {
	repo repository.UserRepository
}

func NewUserGetAllUseCase(repo repository.UserRepository) *UserGetAllUseCase {
	return &UserGetAllUseCase{repo: repo}
}

func (ths *UserGetAllUseCase) Execute(ctx context.Context) (dto.GetAllUsersDTO, error) {
	users, _ := ths.repo.GetAll(ctx)

	usersDTO := dto.GetAllUsersDTO{
		Users: make([]dto.UserOutputDTO, 0, len(users)),
	}

	for _, user := range users {
		usersDTO.Users = append(usersDTO.Users, dto.UserOutputDTO{
			Id:    user.Id,
			Name:  user.Name,
			Login: user.Login,
		})
	}

	return usersDTO, nil
}
