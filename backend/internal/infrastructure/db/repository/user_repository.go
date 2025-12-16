package repository

import (
	"context"
	"log"
	"ritmotrack-backend/internal/domain/entity"
	repoInterface "ritmotrack-backend/internal/domain/repository"
	"ritmotrack-backend/internal/infrastructure/db/model"
	"ritmotrack-backend/internal/infrastructure/db/pg"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
	qb sq.StatementBuilderType
}

func NewUserRepository(db *pg.DB) repoInterface.UserRepository {
	return &UserRepository{
		db: db.GetDB(),
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (ur UserRepository) Create(ctx context.Context, u entity.User) (entity.User, error) {
	query := ur.qb.
		Insert("users").
		Columns("name", "login", "password", "created_at").
		Values(u.Name, u.Login, u.Password, time.Now()).
		Suffix("RETURNING *")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		log.Printf("Error building SQL: %v", err)
		return entity.User{}, err
	}

	var newUser model.User

	err = ur.db.QueryRowxContext(ctx, sqlStr, args...).StructScan(&newUser)
	if err != nil {
		log.Printf("Error executing SQL query: %v\nQuery: %s\nArgs: %v", err, sqlStr, args)
		return entity.User{}, err
	}

	newUserEntity := entity.User{
		Id:        newUser.Id,
		Name:      newUser.Name,
		Login:     newUser.Login,
		Password:  newUser.Password,
		CreatedAt: newUser.CreatedAt,
	}

	return newUserEntity, nil
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
