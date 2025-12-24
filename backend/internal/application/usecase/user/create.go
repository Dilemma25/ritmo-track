package user

import (
	"context"
	"ritmotrack-backend/internal/application/apperror"
	userDTO "ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/provider"
	"ritmotrack-backend/internal/domain/entity"
	"ritmotrack-backend/internal/domain/repository"
)

type CreateUserUseCase interface {
	Execute(ctx context.Context, data userDTO.CreateUserDTO) (*userDTO.UserDTO, error)
}

type createUserUseCase struct {
	repo   repository.UserRepository
	hasher provider.HasherProvider
}

func NewUserCreateUseCase(
	repo repository.UserRepository,
	hasher provider.HasherProvider,
) CreateUserUseCase {
	return &createUserUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (ths *createUserUseCase) Execute(ctx context.Context, dto userDTO.CreateUserDTO) (*userDTO.UserDTO, error) {

	user := entity.NewUser("DefaultName", dto.Login, ths.hasher.Hash(dto.Password))

	existUser, err := ths.repo.GetByLogin(ctx, user.GetLogin())

	if err != nil {
		return nil, err
	}

	if existUser != nil {
		return nil, apperror.ErrUserAlreadyExists
	}

	if err = ths.repo.Store(ctx, user); err != nil {
		return nil, err
	}

	newUserDTO := &userDTO.UserDTO{
		Id:        user.GetId(),
		Name:      user.GetName(),
		Login:     user.GetLogin(),
		CreatedAt: user.GetCreatedAt(),
	}

	return newUserDTO, nil
}
