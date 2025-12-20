package apierror

import (
	"fmt"
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrInternal struct {
	err error
}

func NewErrInternal(err error) *ErrInternal {
	return &ErrInternal{err: err}
}

func (ths *ErrInternal) Error() string {
	return fmt.Errorf("internal server error: %w", ths.err).Error()
}

func (ths *ErrInternal) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "internal_server_error",
		Message: "Внутрення ошибка сервера",
		Detail:  ths.err.Error(),
	}, http.StatusInternalServerError
}
