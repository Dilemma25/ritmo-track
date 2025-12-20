package shema

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}
