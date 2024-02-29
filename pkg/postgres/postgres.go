package postgres

import (
	"fmt"

	"github.com/Snorkin/auth_service/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

//const (
//	maxOpenConns    = 60
//	connMaxLifetime = 120
//)

func CreatePostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	connUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.DBName, cfg.Postgres.Password)
	conn, err := sqlx.Connect(cfg.Postgres.PgDriver, connUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
