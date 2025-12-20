package repository

import (
	"context"
	"errors"
	"ritmotrack-backend/internal/domain/entity"
	repoInterface "ritmotrack-backend/internal/domain/repository"
	"ritmotrack-backend/internal/infrastructure/db/mapper"
	"ritmotrack-backend/internal/infrastructure/db/pg"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pg.MainDB
}

func NewUserRepository(db *pg.MainDB) repoInterface.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ths UserRepository) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	query, args, err := ths.db.GetSq().
		Select("*").
		From("users").
		Where(squirrel.Eq{"login": login}).
		ToSql()

	if err != nil {
		return nil, err
	}

	user, err := mapper.CreateUserFromRow(ths.db.GetPool().QueryRow(ctx, query, args...))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (ths UserRepository) Create(ctx context.Context, user *entity.User) error {
	query, args, err := ths.db.GetSq().
		Insert("users").
		Columns("name", "login", "password", "created_at").
		Values(user.GetName(), user.GetLogin(), user.GetPassword(), user.GetCreatedAt()).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return err
	}

	var id int

	err = ths.db.GetPool().QueryRow(ctx, query, args...).Scan(&id)

	if err != nil {
		return err
	}

	user.SetId(uint(id))

	return nil
}

func (ths UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	users := []*entity.User{
		func() *entity.User { u := entity.NewUser("test1", "test1", ""); return u }(),
		func() *entity.User { u := entity.NewUser("test2", "test2", ""); return u }(),
		func() *entity.User { u := entity.NewUser("test3", "test3", ""); return u }(),
	}

	return users, nil
}

func (ths UserRepository) GetById(ctx context.Context, id uint) (*entity.User, error) {
	return entity.NewUser("test1", "test1", ""), nil
}
