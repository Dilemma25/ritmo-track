package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrAuthFailed struct {
	message string
}

func NewErrAuthFailed(message string) ApiError {
	return &ErrAuthFailed{message: message}
}

func (ths ErrAuthFailed) Error() string {
	return "auth failed"
}

func (ths ErrAuthFailed) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "auth_failed",
		Message: ths.message,
	}, http.StatusUnauthorized
}
