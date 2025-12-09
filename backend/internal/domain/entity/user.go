package entity

import "time"

type User struct {
	Id        int64
	Name      string
	Login     string
	Password  string
	CreatedAt time.Time
}
