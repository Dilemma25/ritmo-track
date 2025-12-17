package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrValidationFailed struct {
	details any
}

func NewErrValidationFailed(errMap map[string]string) *ErrValidationFailed {
	details := make(map[string]string)

	for key, value := range errMap {
		details[key] = value
	}

	return &ErrValidationFailed{
		details: details,
	}
}

func (ths *ErrValidationFailed) Error() string {
	return "validation failed"
}

func (ths *ErrValidationFailed) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "validation_failed",
		Message: "Неверные данные JSON",
		Detail:  ths.details,
	}, http.StatusBadRequest
}
