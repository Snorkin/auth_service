package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   Logger
}

type ServerConfig struct {
	AppVersion  string
	Port        string
	PprofPort   string
	Mode        string // Developer or Production TODO: add enum?
	MaxConnIdle time.Duration
	Timeout     time.Duration
	MaxConnAge  time.Duration
	Time        time.Duration
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

type RedisConfig struct {
	Address      string
	Password     string
	MinIdleConns int
	PoolSize     int
	PoolTimeout  int
	DB           int
}

type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string // Console or Json TODO: add enum?
	Level             string
}

type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

type Session struct {
	Prefix string
	Name   string
	Expire int
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
