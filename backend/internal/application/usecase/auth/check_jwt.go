package auth

import (
	"context"
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/provider"
)

type CheckJwtUseCase interface {
	Execute(ctx context.Context, accessToken string) (*dto.AuthJwtClaimsDTO, error)
}
type checkJwtUseCase struct {
	jwtProvider provider.JwtProvider
}

func NewCheckJwtUseCase(jwtProvider provider.JwtProvider) CheckJwtUseCase {
	return &checkJwtUseCase{jwtProvider: jwtProvider}
}

func (ths *checkJwtUseCase) Execute(ctx context.Context, accessToken string) (*dto.AuthJwtClaimsDTO, error) {
	return ths.jwtProvider.ParseToken(accessToken)
}
