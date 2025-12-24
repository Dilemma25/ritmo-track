package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrInvalidPassword struct{}

func NewErrInvalidPassword() ApiError {
	return &ErrInvalidPassword{}
}

func (ths *ErrInvalidPassword) Error() string {
	return "invalid password"
}

func (ths *ErrInvalidPassword) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "invalid password",
		Message: "неправильный пароль",
	}, http.StatusBadRequest
}
