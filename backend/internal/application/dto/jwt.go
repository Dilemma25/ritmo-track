package dto

type CreateJwtDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JwtOutputDTO struct {
	AccessToken string `json:"access_token"`
}

type AuthJwtClaimsDTO struct {
	UserId uint
}
