package repository

import (
	"context"
	"ritmotrack-backend/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, u entity.User) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	GetById(ctx context.Context, id int64) (entity.User, error)
}
