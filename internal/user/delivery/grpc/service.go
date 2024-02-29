package service

import (
	"github.com/Snorkin/auth_service/config"
	"github.com/Snorkin/auth_service/internal/user"
)

type UsersService struct {
	cfg    *config.Config
	userUC user.Useсase
	//sessUC session.SessionUsecase
}

func NewAuthServiceGRPC(cfg *config.Config, userUC user.Useсase) *UsersService {
	return &UsersService{cfg: cfg, userUC: userUC}
}
