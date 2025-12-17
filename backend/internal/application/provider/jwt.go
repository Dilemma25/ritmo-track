package provider

import "ritmotrack-backend/internal/application/dto"

type JwtProvider interface {
	CreateToken(userId uint) (*dto.JwtDTO, error)
	ParseToken(token string) (*dto.AuthJwtClaimsDTO, error)
	VerifyJwtToken(tokenString string) error
}
