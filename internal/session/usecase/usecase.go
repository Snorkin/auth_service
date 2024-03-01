package usecase

import (
	"context"
	"github.com/Snorkin/auth_service/internal/models"
	"github.com/Snorkin/auth_service/internal/session"
)

type SessionUC struct {
	sessionRepo session.RedisRepository
}

func CreateSessionUC(sessionRepo session.RedisRepository) *SessionUC {
	return &SessionUC{sessionRepo: sessionRepo}
}

func (s *SessionUC) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	return s.sessionRepo.CreateSession(ctx, session, expire)
}

func (s *SessionUC) GetSessionById(ctx context.Context, sessionId string) (*models.Session, error) {
	return s.sessionRepo.GetSessionById(ctx, sessionId)
}

func (s *SessionUC) DeleteById(ctx context.Context, sessionId string) error {
	return s.sessionRepo.DeleteById(ctx, sessionId)
}
