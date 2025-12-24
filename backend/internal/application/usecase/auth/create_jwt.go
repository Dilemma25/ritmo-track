package auth

import (
	"context"
	"ritmotrack-backend/internal/application/apperror"
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/provider"
	"ritmotrack-backend/internal/domain/repository"
)

type CreateJwtUseCase interface {
	Execute(context.Context, dto.CreateJwtDTO) (*dto.JwtDTO, error)
}

type createJwtUseCase struct {
	jwtProvider provider.JwtProvider
	repo        repository.UserRepository
	hasher      provider.HasherProvider
}

func NewCreateJwtUseCase(
	jwtProvider provider.JwtProvider,
	repo repository.UserRepository,
	hasher provider.HasherProvider,
) CreateJwtUseCase {
	return &createJwtUseCase{
		jwtProvider: jwtProvider,
		repo:        repo,
		hasher:      hasher,
	}
}

func (ths *createJwtUseCase) Execute(ctx context.Context, dto dto.CreateJwtDTO) (*dto.JwtDTO, error) {
	user, err := ths.repo.GetByLogin(ctx, dto.Login)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, apperror.ErrUserNotFound
	}

	if !ths.hasher.CompareHashAndPassword(user.GetPassword(), dto.Password) {
		return nil, apperror.ErrInvalidPassword
	}

	tokenDTO, err := ths.jwtProvider.CreateToken(user.GetId())

	if err != nil {
		return nil, err
	}

	return tokenDTO, nil
}
