package config

import "os"

type Config struct {
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func NewConfig() *Config {
	return &Config{
		Port: func() string {

			port := os.Getenv("PORT")
			if port == "" {
				return "8080"
			}

			return port
		}(),

		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		DbName:     os.Getenv("POSTGRES_DB"),
	}
}
