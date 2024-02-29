package service

import (
	"github.com/Snorkin/auth_service/config"
	"github.com/Snorkin/auth_service/internal/user"
	"github.com/Snorkin/auth_service/pkg/logger"
	userService "github.com/Snorkin/auth_service/proto"
)

type UsersService struct {
	logger logger.Logger
	cfg    *config.Config
	userUC user.Useсase
	userService.UnimplementedUserServiceServer
	//sessUC session.SessionUsecase
}

func NewAuthServiceGRPC(logger logger.Logger, cfg *config.Config, userUC user.Useсase) *UsersService {
	return &UsersService{logger: logger, cfg: cfg, userUC: userUC}
}
