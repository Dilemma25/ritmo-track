package responder

import (
	"encoding/json"
	"net/http"
	"ritmotrack-backend/internal/presentation/http/apierror"
)

type Responder struct{}

func NewResponder() *Responder {
	return &Responder{}
}

func (ths *Responder) ResponseOk(w http.ResponseWriter, data any) {
	ths.responseJson(w, data, http.StatusOK)
}

func (ths *Responder) ResponseError(w http.ResponseWriter, error apierror.ApiError) {
	response, status := error.ApiError()

	ths.responseJson(w, response, status)
}

func (ths *Responder) ResponseNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func (ths *Responder) responseJson(w http.ResponseWriter, data any, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		ths.ResponseError(w, apierror.NewErrInternal(err))
	}
}
