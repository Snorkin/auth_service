package main

import (
	"github.com/Snorkin/auth_service/pkg/logger"
	cache "github.com/Snorkin/auth_service/pkg/redis"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
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

	appLogger := logger.CreateLogger(cfg)
	appLogger.Init()
	appLogger.Infof("Version: %s, Mode: %s, Loglevel: %s, SSL: %v", cfg.Server.AppVersion, cfg.Server.Mode, cfg.Logger.Level, cfg.Server.SSL)

	pgDB, err := postgres.CreatePostgresDB(cfg)
	if err != nil {
		log.Fatalf("Cannot connect to pg database, %s\n", err)
	}
	defer func(pgDB *sqlx.DB) {
		err := pgDB.Close()
		if err != nil {
			log.Fatalf("Error while closing db connection, %s", err)
		}
	}(pgDB)

	redisDB := cache.CreateRedisClient(cfg)
	defer func(redisDB *redis.Client) {
		err := redisDB.Close()
		if err != nil {
			log.Fatalf("Error while closing cache db connection, %s", err)
		}
	}(redisDB)

	server := authServer.CreateAuthApp(appLogger, cfg, pgDB, redisDB)
	appLogger.Fatal(server.Run())
}
