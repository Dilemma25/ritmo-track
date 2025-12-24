package controller

import (
	"errors"
	"net/http"
	"ritmotrack-backend/internal/application/apperror"
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/application/usecase/auth"
	"ritmotrack-backend/internal/presentation/http/apierror"
	"ritmotrack-backend/internal/presentation/http/responder"
	"ritmotrack-backend/internal/presentation/http/shema"
	"ritmotrack-backend/internal/presentation/http/validator"
)

type AuthController struct {
	createJwtUseCase auth.CreateJwtUseCase
	checkJwtUseCase  auth.CheckJwtUseCase
	validator        *validator.Validator
	responder        *responder.Responder
}

func NewAuthController(
	createJwtUseCase auth.CreateJwtUseCase,
	checkJwtUseCase auth.CheckJwtUseCase,
	validator *validator.Validator,
	responder *responder.Responder,
) *AuthController {
	return &AuthController{
		createJwtUseCase: createJwtUseCase,
		checkJwtUseCase:  checkJwtUseCase,
		validator:        validator,
		responder:        responder,
	}
}

func (ths *AuthController) CreateJwt(w http.ResponseWriter, r *http.Request) {
	var req *shema.AuthRequest

	if err := ths.validator.ValidateBody(r, &req); err != nil {
		ths.responder.ResponseError(w, err)

		return
	}

	jwtDto := dto.CreateJwtDTO{
		Login:    req.Login,
		Password: req.Password,
	}

	token, err := ths.createJwtUseCase.Execute(r.Context(), jwtDto)

	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			ths.responder.ResponseError(w, apierror.NewErrUserNotFound())
		}

		if errors.Is(err, apperror.ErrInvalidPassword) {
			ths.responder.ResponseError(w, apierror.NewErrInvalidPassword())
		}

		ths.responder.ResponseError(w, apierror.NewErrInternal(err))

		return
	}

	tokenResponse := shema.AuthResponse{
		Access: token.AccessToken,
	}

	ths.responder.ResponseOk(w, tokenResponse)
}

func (ths *AuthController) CheckJwt(w http.ResponseWriter, r *http.Request) {
	ths.responder.ResponseNoContent(w)
}
