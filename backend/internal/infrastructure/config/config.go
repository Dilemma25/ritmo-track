package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	JwtSecretKey string
	JwtTTL       int
}

func NewConfig() *Config {

	jwtTTL, err := strconv.Atoi(os.Getenv("JWT_TTL"))

	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		Port: os.Getenv("PORT"),

		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		DbName:     os.Getenv("POSTGRES_DB"),

		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
		JwtTTL:       jwtTTL,
	}
}

func (ths *Config) DbDNS() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		ths.DbHost, ths.DbUser, ths.DbPassword, ths.DbName, ths.DbPort,
	)
}
