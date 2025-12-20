package provider

import (
	"ritmotrack-backend/internal/application/provider"

	"golang.org/x/crypto/bcrypt"
)

type hasherProvider struct{}

func NewHasherProvider() provider.HasherProvider {
	return &hasherProvider{}
}

func (ths hasherProvider) Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (ths hasherProvider) CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
