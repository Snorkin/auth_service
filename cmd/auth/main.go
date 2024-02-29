package main

import (
	"log"

	"github.com/Snorkin/auth_service/config"
	authServer "github.com/Snorkin/auth_service/internal/app"
	"github.com/Snorkin/auth_service/pkg/postgres"
)

func main() {
	log.Println("Starting auth service")

	cfg, err := config.GetConfig(".env")
	if err != nil {
		log.Fatalln("Cannot load env variables")
	}

	pgDB, err := postgres.CreatePostgresDB(cfg)
	if err != nil {
		log.Fatalf("Cannot connect to pg database, %s\n", err)
	}

	server := authServer.CreateAuthApp(cfg, pgDB)
	err = server.Run()
	if err != nil {
		log.Fatalf("Cannot run the app")
	}
}
