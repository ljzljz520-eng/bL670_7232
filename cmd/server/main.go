package main

import (
	"log"
	"net/http"

	"training-review/internal/api"
	"training-review/internal/config"
	"training-review/internal/service"
	"training-review/internal/store"
)

func main() {
	cfg := config.FromEnv()
	if !cfg.Validate() {
		log.Fatal("invalid configuration")
	}
	database, err := store.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	clock := service.FixedClock{Value: "2000-01-01T00:00:00Z"}
	registrations := service.NewRegistrationService(database, clock)
	workflows := service.NewWorkflowService(database, clock)
	archive := service.NewArchiveService(database, clock)
	server := api.NewServer(registrations, workflows, archive)
	log.Printf("training review listening on %s", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, server.Handler()); err != nil {
		log.Fatal(err)
	}
}
