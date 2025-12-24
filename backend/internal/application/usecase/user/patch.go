package user

import (
	"context"
	"ritmotrack-backend/internal/application/apperror"
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/provider"
	"ritmotrack-backend/internal/domain/repository"
)

type PatchUserUseCase interface {
	Execute(ctx context.Context, userID uint, dto dto.PatchUserDTO) error
}

type patchUserUseCase struct {
	repo   repository.UserRepository
	hasher provider.HasherProvider
}

func NewPatchUserUseCase(
	repo repository.UserRepository,
	hasher provider.HasherProvider,
) PatchUserUseCase {
	return &patchUserUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (ths *patchUserUseCase) Execute(ctx context.Context, userId uint, dto dto.PatchUserDTO) error {
	user, err := ths.repo.GetById(ctx, userId)

	if err != nil {
		return err
	}

	if user == nil {
		return apperror.ErrUserNotFound
	}

	if dto.Name != nil {
		if domainErr := user.SetName(*dto.Name); domainErr != nil {
			return domainErr
		}
	}

	if dto.Password != nil {
		newHashedPassword := ths.hasher.Hash(*dto.Password)

		if domainErr := user.SetPassword(newHashedPassword); domainErr != nil {
			return domainErr
		}
	}

	err = ths.repo.Update(ctx, user)

	return err
}
