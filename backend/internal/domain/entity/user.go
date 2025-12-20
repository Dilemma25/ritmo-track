package entity

import "time"

type User struct {
	id        uint
	name      string
	login     string
	password  string
	createdAt time.Time
}

func NewUser(name string, login string, password string) *User {
	user := &User{
		name:      name,
		login:     login,
		password:  password,
		createdAt: time.Now(),
	}

	return user
}

func (ths *User) GetId() uint {
	return ths.id
}

func (ths *User) SetId(id uint) {
	ths.id = id
}

func (ths *User) GetName() string {
	return ths.name
}

func (ths *User) SetName(name string) {
	ths.name = name
}

func (ths *User) GetLogin() string {
	return ths.login
}

func (ths *User) SetLogin(login string) {
	ths.login = login
}

func (ths *User) GetPassword() string {
	return ths.password
}

func (ths *User) SetPassword(password string) {
	ths.password = password
}

func (ths *User) GetCreatedAt() time.Time {
	return ths.createdAt
}

func (ths *User) SetCreatedAt(createdAt time.Time) {
	ths.createdAt = createdAt
}
