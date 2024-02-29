package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	AppVersion string
	Port       string
	PprofPort  string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
	PgDriver string
}

func GetConfig(path string) (*Config, error) {
	err := loadConfig(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	cfg.Postgres.Port = os.Getenv("DB_PORT")
	cfg.Postgres.Host = os.Getenv("DB_HOST")
	cfg.Postgres.User = os.Getenv("DB_USER")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.DBName = os.Getenv("DB_NAME")
	cfg.Postgres.SSLMode, _ = strconv.ParseBool(os.Getenv("DB_SSL"))
	cfg.Postgres.PgDriver = os.Getenv("DB_DRIVER")

	return &cfg, nil
}

func loadConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
