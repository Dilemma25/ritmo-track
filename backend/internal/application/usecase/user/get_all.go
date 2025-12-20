package user

import (
	"context"
	userDTO "ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/domain/repository"
)

type GetAllUsersUseCase interface {
	Execute(ctx context.Context) ([]*userDTO.UserDTO, error)
}

type getAllUsersUseCase struct {
	repo repository.UserRepository
}

func NewUserGetAllUseCase(repo repository.UserRepository) GetAllUsersUseCase {
	return &getAllUsersUseCase{repo: repo}
}

func (ths *getAllUsersUseCase) Execute(ctx context.Context) ([]*userDTO.UserDTO, error) {
	users, _ := ths.repo.GetAll(ctx)

	var usersDTO []*userDTO.UserDTO

	for _, user := range users {
		usersDTO = append(usersDTO, &userDTO.UserDTO{
			Id:    user.GetId(),
			Name:  user.GetName(),
			Login: user.GetLogin(),
		})
	}

	return usersDTO, nil
}
