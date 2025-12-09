package repository

import (
	"context"
	"ritmotrack-backend/internal/domain/entity"
	"time"
)

type UserRepository struct{}

func (ur UserRepository) Create(ctx context.Context, u entity.User) (entity.User, error) {
	newUser := entity.User{
		Id:        1,
		Name:      u.Name,
		Login:     u.Login,
		Password:  u.Password,
		CreatedAt: time.Now(),
	}

	return newUser, nil
}

func (ur UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{
		{Id: 1, Name: "test1", Login: "test1"},
		{Id: 2, Name: "test2", Login: "test2"},
		{Id: 3, Name: "test3", Login: "test3"},
	}

	return users, nil
}

func (ur UserRepository) GetById(ctx context.Context, id int64) (entity.User, error) {
	return entity.User{
		Id:        id,
		Name:      "test",
		Login:     "test",
		Password:  "****",
		CreatedAt: time.Now(),
	}, nil
}
