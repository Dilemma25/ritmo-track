package auth

import (
	"context"
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/provider"
	"ritmotrack-backend/internal/domain/entity"
	"ritmotrack-backend/internal/domain/repository"
)

type CreateJwtUseCase interface {
	Execute(context.Context, dto.CreateJwtDTO) (*dto.JwtDTO, error)
}

type createJwtUseCase struct {
	jwtProvider provider.JwtProvider
	repo        repository.UserRepository
}

func NewCreateJwtUseCase(jwtProvider provider.JwtProvider, repo repository.UserRepository) CreateJwtUseCase {
	return &createJwtUseCase{jwtProvider: jwtProvider, repo: repo}
}

func (ths *createJwtUseCase) Execute(ctx context.Context, data dto.CreateJwtDTO) (*dto.JwtDTO, error) {
	user, err := ths.repo.GetByLogin(ctx, data.Login)

	if err != nil {
		return nil, err
	}

	if user == nil {
		user = entity.NewUser("test", data.Login, data.Password)

		err = ths.repo.Create(ctx, user)

		if err != nil {
			return nil, err
		}
	}

	tokenDTO, err := ths.jwtProvider.CreateToken(user.GetId())

	if err != nil {
		return nil, err
	}

	return tokenDTO, nil
}
