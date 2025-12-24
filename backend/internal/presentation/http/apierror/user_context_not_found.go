package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrUserContextNotFound struct{}

func NewErrUserContextNotFound() ApiError {
	return &ErrUserContextNotFound{}
}

func (ths *ErrUserContextNotFound) Error() string {
	return "user context not found"
}

func (ths *ErrUserContextNotFound) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "user_context_not_found",
		Message: "контекст пользователя не найден",
	}, http.StatusBadRequest
}
