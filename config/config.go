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
	SSLMode  string
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
	//TODO: check viper
	cfg := Config{}

	cfg.Server.AppVersion = os.Getenv("SERVER_VERSION")
	cfg.Server.Port = os.Getenv("SERVER_PORT")
	cfg.Server.Mode = os.Getenv("SERVER_MODE")
	cfg.Server.MaxConnIdle, _ = time.ParseDuration(os.Getenv("SERVER_MAXCONNIDLE"))
	cfg.Server.Timeout, _ = time.ParseDuration(os.Getenv("SERVER_TIMEOUT"))
	cfg.Server.MaxConnAge, _ = time.ParseDuration(os.Getenv("SERVER_MAXCONNAGE"))
	cfg.Server.Time, _ = time.ParseDuration(os.Getenv("SERVER_TIME"))

	cfg.Postgres.Port = os.Getenv("DB_PORT")
	cfg.Postgres.Host = os.Getenv("DB_HOST")
	cfg.Postgres.User = os.Getenv("DB_USER")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.DBName = os.Getenv("DB_NAME")
	cfg.Postgres.SSLMode = os.Getenv("DB_SSL")
	cfg.Postgres.PgDriver = os.Getenv("DB_DRIVER")

	cfg.Redis.DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	cfg.Redis.Address = os.Getenv("REDIS_ADDRESS")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
	cfg.Redis.PoolTimeout, _ = strconv.Atoi(os.Getenv("REDIS_POOLTIMEOUT"))
	cfg.Redis.PoolSize, _ = strconv.Atoi(os.Getenv("REDIS_POOLSIZE"))
	cfg.Redis.MinIdleConns, _ = strconv.Atoi(os.Getenv("REDIS_MINIDLECONN"))

	check := false
	if os.Getenv("SERVER_MODE") == "Development" {
		check = true
	}
	cfg.Logger.Development = check
	cfg.Logger.Encoding = os.Getenv("LOGGER_ENCODING")

	return &cfg, nil
}

func loadConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
