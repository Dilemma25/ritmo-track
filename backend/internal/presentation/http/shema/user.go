package shema

import "time"

type CreateUserRequest struct {
	Login    string `json:"login" validate:"required,min=5,max=30,regex=^[a-zA-Z0-9_]+$"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type PatchUserRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=1,max=30"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=8,max=20"`
}

type UserResponse struct {
	Id        uint      `json:"id"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"createdAt"`
}
