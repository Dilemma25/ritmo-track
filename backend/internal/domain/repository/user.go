package repository

import (
	"context"
	"ritmotrack-backend/internal/domain/entity"
)

type UserRepository interface {
	Store(ctx context.Context, user *entity.User) error
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetById(ctx context.Context, id uint) (*entity.User, error)
	GetByLogin(ctx context.Context, login string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}
