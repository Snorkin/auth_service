package app

import (
	"github.com/Snorkin/auth_service/config"
	sessRepository "github.com/Snorkin/auth_service/internal/session/repository"
	sessUsecase "github.com/Snorkin/auth_service/internal/session/usecase"
	userHandler "github.com/Snorkin/auth_service/internal/user/delivery/grpc"
	userRepository "github.com/Snorkin/auth_service/internal/user/repository"
	userUsecase "github.com/Snorkin/auth_service/internal/user/usecase"
	"github.com/Snorkin/auth_service/pkg/interceptor"
	"github.com/Snorkin/auth_service/pkg/logger"
	grpcService "github.com/Snorkin/auth_service/proto"
	"github.com/go-redis/redis/v8"
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
	redis  *redis.Client
}

func CreateAuthApp(logger logger.Logger, cfg *config.Config, db *sqlx.DB, redis *redis.Client) *App {
	return &App{logger: logger, cfg: cfg, db: db, redis: redis}
}

func (a *App) Run() error {
	ic := interceptor.CreateInterceptor(a.logger, a.cfg)
	userRepo := userRepository.CreateUserPgRepo(a.db)
	sessRepo := sessRepository.CreateSessionRepository(a.redis)
	userRedisRepo := userRepository.CreateUserRedisRepository(a.redis)
	userUC := userUsecase.CreateUserUseCase(a.logger, userRepo, userRedisRepo)
	sessUC := sessUsecase.CreateSessionUC(sessRepo)

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: a.cfg.Server.MaxConnIdle * time.Minute,
			Timeout:           a.cfg.Server.Timeout * time.Second,
			MaxConnectionAge:  a.cfg.Server.MaxConnAge * time.Minute,
			Time:              a.cfg.Server.Time * time.Minute,
		}),
		grpc.UnaryInterceptor(ic.Log))
	if a.cfg.Server.Mode != "Production" {
		reflection.Register(grpcServer)
	}

	authHandlerGRPC := userHandler.CreateAuthHandlerGRPC(a.logger, a.cfg, userUC, sessUC)
	grpcService.RegisterUserServiceServer(grpcServer, authHandlerGRPC)

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
