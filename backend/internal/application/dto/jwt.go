package dto

type CreateJwtDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JwtDTO struct {
	AccessToken string `json:"access_token"`
}

type AuthJwtClaimsDTO struct {
	UserId uint `json:"user_id"`
}
