package main

import (
	"log"
	"prs/internal/config"
	"prs/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// DB

    cfg := config.LoadConfig()

    db, err:= gorm.Open(postgres.Open(cfg.Postgres.DSN()), &gorm.Config{})
    if err != nil {
        log.Fatalf("db connect error: %v", err)
    }

    repo := repository.NewRepository(db)


	
}