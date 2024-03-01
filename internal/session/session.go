package session

import (
	"context"
	"github.com/Snorkin/auth_service/internal/models"
)

type UseCase interface {
	CreateSession(ctx context.Context, session *models.Session, expire int) (string, error)
	GetSessionById(ctx context.Context, sessionId string) (*models.Session, error)
	DeleteById(ctx context.Context, sessionId string) error
}

type RedisRepository interface {
	CreateSession(ctx context.Context, session *models.Session, expire int) (string, error)
	GetSessionById(ctx context.Context, sessionId string) (*models.Session, error)
	DeleteById(ctx context.Context, sessionId string) error
}
