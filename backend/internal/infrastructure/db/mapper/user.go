package mapper

import (
	"ritmotrack-backend/internal/domain/entity"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateUserFromRow(row pgx.Row) (*entity.User, error) {
	var id uint
	var name string
	var login string
	var password string
	var createdAt time.Time

	err := row.Scan(&id, &name, &login, &password, &createdAt)

	if err != nil {
		return nil, err
	}

	user := &entity.User{}

	user.SetId(id)
	user.SetName(name)
	user.SetLogin(login)
	user.SetPassword(password)
	user.SetCreatedAt(createdAt)

	return user, nil
}
