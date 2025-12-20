package provider

import (
	"fmt"
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/provider"
	"ritmotrack-backend/internal/infrastructure/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtClaims struct {
	Sub uint `json:"sub"`
	jwt.RegisteredClaims
}
type jwtProvider struct {
	config *config.Config
}

func NewJwtProvider(config *config.Config) provider.JwtProvider {
	return &jwtProvider{
		config: config,
	}
}

func (ths *jwtProvider) CreateToken(userId uint) (*dto.JwtDTO, error) {
	expiresAt := time.Now().Add(time.Duration(ths.config.JwtTTL) * time.Second)

	claims := jwtClaims{
		Sub: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(ths.config.JwtSecretKey))

	if err != nil {
		return nil, err
	}

	return &dto.JwtDTO{
		AccessToken: tokenString,
	}, nil
}

func (ths *jwtProvider) ParseToken(accessToken string) (*dto.AuthJwtClaimsDTO, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&jwtClaims{},
		ths.getSecretKey,
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return &dto.AuthJwtClaimsDTO{
		UserId: claims.Sub,
	}, nil

}

func (ths *jwtProvider) getSecretKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return []byte(ths.config.JwtSecretKey), nil
}

func (ths *jwtProvider) VerifyJwtToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
