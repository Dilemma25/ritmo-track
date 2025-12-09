package config

import "os"

type Config struct {
	Port string
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
	}
}
