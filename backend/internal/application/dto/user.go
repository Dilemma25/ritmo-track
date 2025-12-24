package dto

import "time"

type CreateUserDTO struct {
	Name     string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PatchUserDTO struct {
	Name     *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserDTO struct {
	Id        uint      `json:"id"`
	Name      string    `json:"username"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"createdAt"`
}
