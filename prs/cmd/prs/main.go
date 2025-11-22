package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"prs/internal/config"
	"prs/internal/handler"
	"prs/internal/repository"
	"syscall"

	"time"

	"github.com/go-chi/chi/v5"
)

func main() {

	// Initialize DB
    cfg := config.LoadConfig()
    repo, err := repository.NewRepository(cfg.Postgres.DSN())
	if err != nil {
		log.Fatalf("Database error %v", err)
	}


	// TODO: create service

	//  Initialize router

	r := chi.NewRouter()

	handler.RegisterRouters(r) // TODO: mapping 

	server := &http.Server{
		Addr:         ":" + cfg.HTTP.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start service

	sigTermChan := make(chan os.Signal, 1)
	signal.Notify(sigTermChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed")
		}
	}()

	log.Println("PRs Service started")

	// Shutdown by signal

	<-sigTermChan

	log.Println("Stopping PRs Service...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("PRs graceful shutdown failed: %s", err)
	}

	log.Println("PRs service gracefully stopped.")
}