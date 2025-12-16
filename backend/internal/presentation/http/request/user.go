package request

type UserCreateRequest struct {
	Name     string `validate:"required,min=5,max=30,regex=^[a-zA-Z0-9_]+$"`
	Login    string `validate:"required,min=5,max=30,regex=^[a-zA-Z0-9_]+$"`
	Password string `validate:"required,min=8,max=20"`
}
