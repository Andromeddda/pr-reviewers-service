package config

import (
	"fmt"
	"log"
	"os"
)

type PostgresConfig struct {
	Host string
	Port string 
	User string 
	Pass string
	DB 	 string 
}

type HTTPConfig struct {
	Port string
}

type Config struct {
	Postgres 	PostgresConfig
	HTTP 		HTTPConfig
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
		HTTP: HTTPConfig {
			Port: getEnvOrDefault("HTTP_PORT", "8080"),
		},
	}

	return cfg
}


func (cfg PostgresConfig) DSN() string {
	dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.Host, cfg.User, cfg.Pass, cfg.DB, cfg.Port,
    )

	log.Printf("DSN: %s", dsn)

	return dsn
}

func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	log.Printf("Key %s not found. Using default %s", key, def)
	return def
}