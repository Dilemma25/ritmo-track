package mapper

import (
	"ritmotrack-backend/internal/application/dto"
	"ritmotrack-backend/internal/presentation/http/shema"
)

func CreateUserResponseFromDTO(dto *dto.UserDTO) shema.UserResponse {
	return shema.UserResponse{
		Id:        dto.Id,
		Login:     dto.Login,
		CreatedAt: dto.CreatedAt,
	}
}
