package dto

import "time"

type CreateUserDTO struct {
	Name     string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserOutputDTO struct {
	Id        uint      `json:"id"`
	Name      string    `json:"username"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllUsersDTO struct {
	Users []UserOutputDTO
}
