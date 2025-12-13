package repository

import (
	"context"
	"database/sql"
	"log"
	"ritmotrack-backend/internal/domain/entity"
	repoInterface "ritmotrack-backend/internal/domain/repository"
	"ritmotrack-backend/internal/infrastructure/db/pg"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db *sql.DB
	qb sq.StatementBuilderType
}

func NewUserRepository(db *pg.DB) repoInterface.UserRepository {
	return &UserRepository{
		db: db.GetDB(),
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (ur UserRepository) Create(ctx context.Context, u entity.User) (entity.User, error) {
	// Формируем SQL INSERT с RETURNING для получения ID и CreatedAt
	query := ur.qb.
		Insert("users"). // имя таблицы в БД
		Columns("name", "login", "password", "created_at").
		Values(u.Name, u.Login, u.Password, time.Now()).
		Suffix("RETURNING id, created_at") // возвращаем ID и дату создания

	// Генерируем SQL и аргументы
	sqlStr, args, err := query.ToSql()
	if err != nil {
		log.Printf("Error building SQL: %v", err)
		return entity.User{}, err
	}

	var id int64
	var createdAt time.Time

	// Выполняем запрос
	err = ur.db.QueryRowContext(ctx, sqlStr, args...).Scan(&id, &createdAt)
	if err != nil {
		log.Printf("Error executing SQL query: %v\nQuery: %s\nArgs: %v", err, sqlStr, args)
		return entity.User{}, err
	}

	// Возвращаем созданного пользователя с реальными данными из БД
	newUser := entity.User{
		Id:        id,
		Name:      u.Name,
		Login:     u.Login,
		Password:  u.Password, // можно хранить хэш, если есть
		CreatedAt: createdAt,
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
