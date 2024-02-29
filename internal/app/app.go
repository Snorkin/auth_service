package app

import (
	"github.com/Snorkin/auth_service/config"
	userService "github.com/Snorkin/auth_service/internal/user/delivery/grpc"
	userRepository "github.com/Snorkin/auth_service/internal/user/repository"
	userUsecase "github.com/Snorkin/auth_service/internal/user/usecase"
	"github.com/Snorkin/auth_service/pkg/interceptor"
	"github.com/Snorkin/auth_service/pkg/logger"
	grpcService "github.com/Snorkin/auth_service/proto"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	logger logger.Logger
	cfg    *config.Config
	db     *sqlx.DB
}

func CreateAuthApp(cfg *config.Config, db *sqlx.DB) *App {
	return &App{cfg: cfg, db: db}
}

func (a *App) Run() error {
	ic := interceptor.CreateInterceptor(a.logger, a.cfg)
	userRepo := userRepository.NewUserPgRepo(a.db)
	userUC := userUsecase.CreateUserUseCase(userRepo)

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: a.cfg.Server.MaxConnIdle * time.Minute,
			Timeout:           a.cfg.Server.Timeout * time.Second,
			MaxConnectionAge:  a.cfg.Server.MaxConnAge * time.Minute,
			Time:              a.cfg.Server.Time * time.Minute,
		}),
		grpc.UnaryInterceptor(ic.Log))
	if a.cfg.Server.Mode == "Production" {
		reflection.Register(grpcServer)
	}

	authServiceGRPC := userService.NewAuthServiceGRPC(a.logger, a.cfg, userUC)
	grpcService.RegisterUserServiceServer(grpcServer, authServiceGRPC)

	listener, err := net.Listen("tcp", a.cfg.Server.Port)
	if err != nil {
		return err
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			a.logger.Fatalf("Error: closing tcp connection, %s", err)
		}
	}(listener)

	go func() {
		a.logger.Infof("Server is listening on port: %v", a.cfg.Server.Port)
		if err := grpcServer.Serve(listener); err != nil {
			a.logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	grpcServer.GracefulStop()
	a.logger.Info("Server exited properly")

	return nil
}
