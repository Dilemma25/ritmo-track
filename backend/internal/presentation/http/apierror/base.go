package apierror

import "ritmotrack-backend/internal/presentation/http/shema"

type ApiError interface {
	ApiError() (*shema.ErrorResponse, int)
}
