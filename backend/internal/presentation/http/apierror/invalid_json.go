package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrInvalidJson struct {
	err error
}

func NewErrInvalidJson(err error) ApiError {
	return &ErrInvalidJson{err: err}
}

func (ths *ErrInvalidJson) Error() string {
	return "invalid json"
}

func (ths *ErrInvalidJson) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "invalid_json",
		Message: "Неверный формат JSON",
		Detail:  ths.err.Error(),
	}, http.StatusBadRequest
}
