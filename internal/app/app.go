package app

import (
	"fmt"
	"github.com/Snorkin/auth_service/config"
	userService "github.com/Snorkin/auth_service/internal/user/delivery/grpc"
	userRepository "github.com/Snorkin/auth_service/internal/user/repository"
	userUsecase "github.com/Snorkin/auth_service/internal/user/usecase"
	"github.com/jmoiron/sqlx"
)

type App struct {
	// logger logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func CreateAuthApp(cfg *config.Config, db *sqlx.DB) *App {
	return &App{cfg: cfg, db: db}
}

func (a *App) Run() error {
	userRepo := userRepository.NewUserPgRepo(a.db)
	userUC := userUsecase.CreateUserUseCase(userRepo)

	authGRPCServer := userService.NewAuthServiceGRPC(a.cfg, userUC)
	fmt.Println(authGRPCServer)
	return nil
}
