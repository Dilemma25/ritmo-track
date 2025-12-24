package apierror

import (
	"net/http"
	"ritmotrack-backend/internal/presentation/http/shema"
)

type ErrInvalidUrlParameter struct {
	err error
}

func NewErrInvalidUrlParameter(err error) ApiError {
	return &ErrInvalidUrlParameter{err: err}
}

func (ths ErrInvalidUrlParameter) Error() string {
	return "invalid url parameter"
}

func (ths ErrInvalidUrlParameter) ApiError() (*shema.ErrorResponse, int) {
	return &shema.ErrorResponse{
		Error:   "invalid_url_parameter",
		Message: "неверный формат URL параметра",
		Detail:  ths.err.Error(),
	}, http.StatusBadRequest
}
