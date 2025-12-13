package dto

type CreateUserDTO struct {
	Name     string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserOutputDTO struct {
	Id    int64  `json:"id"`
	Name  string `json:"username"`
	Login string `json:"login"`
}

type GetAllUsersDTO struct {
	Users []UserOutputDTO
}
