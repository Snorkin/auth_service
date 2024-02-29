package handlers

import (
	"github.com/Snorkin/auth_service/config"
	"github.com/Snorkin/auth_service/internal/user"
	"github.com/Snorkin/auth_service/pkg/logger"
	userService "github.com/Snorkin/auth_service/proto"
)

type UsersHandler struct {
	logger logger.Logger
	cfg    *config.Config
	userUC user.Useсase
	userService.UnimplementedUserServiceServer
	//sessUC session.SessionUsecase
}

func NewAuthHandlerGRPC(logger logger.Logger, cfg *config.Config, userUC user.Useсase) *UsersHandler {
	return &UsersHandler{logger: logger, cfg: cfg, userUC: userUC}
}
