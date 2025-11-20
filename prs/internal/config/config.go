package config

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	Host string
	Port string 
	User string 
	Pass string
	DB 	 string 
}

type Config struct {
	Postgres PostgresConfig
}

func LoadConfig() *Config {
	cfg := &Config{
		Postgres: PostgresConfig{
			Host: getEnvOrDefault("POSTGRES_HOST", "localhost"),
			Port: getEnvOrDefault("POSTGRES_PORT", "5432"),
			User: getEnvOrDefault("POSTGRES_USER", "postgres"),
			Pass: getEnvOrDefault("POSTGRES_PASS", "postgres"),
			DB:   getEnvOrDefault("POSTGRES_DB",   "prs"),
		},
	}

	return cfg
}


func (cfg PostgresConfig) DSN() string {
	return fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.Host, cfg.User, cfg.Pass, cfg.DB, cfg.Port,
    )
}

func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v != "" {
		return v 
	}
	return def
}