package usecase

import (
	"context"
	"ritmotrack-backend/internal/application/user/dto"
	"ritmotrack-backend/internal/domain/repository"
)

type UserGetAllUseCase struct {
	Repo repository.UserRepository
}

func (ths *UserGetAllUseCase) Execute(ctx context.Context) (dto.GetAllUsersDTO, error) {
	users, _ := ths.Repo.GetAll(ctx)

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
