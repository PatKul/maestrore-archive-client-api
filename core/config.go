package core

import (
	"os"
)

type Config struct {
	Port             string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
}

func NewConfig() *Config {
	return &Config{
		Port:             getValueOrDefault(os.Getenv("PORT"), ":8080"),
		DatabaseHost:     getValueOrDefault(os.Getenv("DATABASE_HOST"), "localhost"),
		DatabaseUser:     getValueOrDefault(os.Getenv("DATABASE_USER"), "root"),
		DatabasePassword: getValueOrDefault(os.Getenv("DATABASE_PASSWORD"), "admin@2005"),
	}
}

func getValueOrDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}
