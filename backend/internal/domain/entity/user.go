package entity

import (
	"errors"
	"strings"
	"time"
)

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

func (ths *User) SetName(name string) error {
	//TODO сделать ошибки и валидацию для домена
	if strings.TrimSpace(name) == "" {
		return errors.New("недопустимое имя")
	}

	ths.name = name
	return nil
}

func (ths *User) GetLogin() string {
	return ths.login
}

func (ths *User) SetLogin(login string) error {
	//TODO сделать ошибки и валидацию для домена
	if strings.TrimSpace(login) == "" {
		return errors.New("недопустимый логин")
	}

	ths.login = login
	return nil
}

func (ths *User) GetPassword() string {
	return ths.password
}

func (ths *User) SetPassword(password string) error {
	//TODO сделать ошибки и валидацию для домена
	if strings.TrimSpace(password) == "" {
		return errors.New("недопустимый пароль")
	}

	ths.password = password
	return nil
}

func (ths *User) GetCreatedAt() time.Time {
	return ths.createdAt
}

func (ths *User) SetCreatedAt(createdAt time.Time) {
	ths.createdAt = createdAt
}
