package config

import (
	"fmt"
	"os"
)

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSLMode  string
}

func Load() Config {
	return Config{
		Postgres: PostgresConfig{
			Host:     getEnv("PG_HOST"),
			Port:     getEnv("PG_PORT"),
			User:     getEnv("PG_USER"),
			Password: getEnv("PG_PASSWORD"),
			DB:       getEnv("PG_DB"),
			SSLMode:  getEnv("PG_SSLMODE"),
		},
	}
}

func (p PostgresConfig) ConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DB,
		p.SSLMode,
	)
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("missing env: " + key)
	}

	return value
}
