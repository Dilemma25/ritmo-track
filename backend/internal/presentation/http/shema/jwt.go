package shema

type AuthRequest struct {
	Login    string `json:"login" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}

type AuthResponse struct {
	Access string `json:"accessToken"`
}
