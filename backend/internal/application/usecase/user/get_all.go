package user

import (
	"context"
	userDTO "ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/domain/repository"
)

type GetAllUsersUseCase struct {
	repo repository.UserRepository
}

func NewUserGetAllUseCase(repo repository.UserRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{repo: repo}
}

func (ths *GetAllUsersUseCase) Execute(ctx context.Context) (userDTO.GetAllUsersDTO, error) {
	users, _ := ths.repo.GetAll(ctx)

	usersDTO := userDTO.GetAllUsersDTO{
		Users: make([]userDTO.UserOutputDTO, 0, len(users)),
	}

	for _, user := range users {
		usersDTO.Users = append(usersDTO.Users, userDTO.UserOutputDTO{
			Id:    user.GetId(),
			Name:  user.GetName(),
			Login: user.GetLogin(),
		})
	}

	return usersDTO, nil
}
