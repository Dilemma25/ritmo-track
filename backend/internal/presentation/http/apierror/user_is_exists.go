package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrUserAlreadyExists struct{}

func NewErrUserAlreadyExists() ApiError {
	return &ErrUserAlreadyExists{}
}

func (ths *ErrUserAlreadyExists) Error() string {
	return "user already exists"
}

func (ths *ErrUserAlreadyExists) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "user_already_exists",
		Message: "Пользователь с таким логином уже существует",
	}, http.StatusBadRequest
}
