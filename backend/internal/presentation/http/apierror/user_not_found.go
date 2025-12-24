package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrUserNotFound struct{}

func NewErrUserNotFound() ApiError {
	return &ErrUserNotFound{}
}

func (err *ErrUserNotFound) Error() string {
	return "user not found"
}

func (err *ErrUserNotFound) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "user_not_found",
		Message: "Пользователь не найдет",
	}, http.StatusNotFound
}
