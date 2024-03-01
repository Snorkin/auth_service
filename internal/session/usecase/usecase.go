package usecase

import (
	"context"
	"github.com/Snorkin/auth_service/config"
	"github.com/Snorkin/auth_service/internal/models"
	"github.com/Snorkin/auth_service/internal/session"
)

type sessionUC struct {
	sessionRepo session.RedisRepository
	cfg         *config.Config
}

func CreateSessionUC(sessionRepo session.RedisRepository, cfg *config.Config) session.UseCase {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg}
}

func (s *sessionUC) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	return s.sessionRepo.CreateSession(ctx, session, expire)
}

func (s *sessionUC) GetSessionById(ctx context.Context, sessionId string) (*models.Session, error) {
	return s.sessionRepo.GetSessionById(ctx, sessionId)
}

func (s *sessionUC) DeleteById(ctx context.Context, sessionId string) error {
	return s.sessionRepo.DeleteById(ctx, sessionId)
}
