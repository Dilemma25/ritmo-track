package model

import "time"

type User struct {
	Id        int64     `db:"id"`
	Login     string    `db:"login"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
