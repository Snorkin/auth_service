package server

import (
	"fmt"
	"github.com/Snorkin/auth_service/config"
	grpcServer "github.com/Snorkin/auth_service/internal/user/delivery/grpc"
	userRepository "github.com/Snorkin/auth_service/internal/user/repository"
	userUsecase "github.com/Snorkin/auth_service/internal/user/usecase"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	// logger logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func CreateAuthServer(cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{cfg: cfg, db: db}
}

func (a *Server) Run() error {
	userRepo := userRepository.NewUserPgRepo(a.db)
	userUC := userUsecase.CreateUserUseCase(userRepo)

	authGRPCServer := grpcServer.NewAuthServerGRPC(a.cfg, userUC)
	fmt.Println(authGRPCServer)
	return nil
}
