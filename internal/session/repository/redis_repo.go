package repository

import (
	"context"
	"encoding/json"
	"github.com/Snorkin/auth_service/internal/models"
	"github.com/Snorkin/auth_service/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

const (
	prefix = "session:"
)

type SessionRepo struct {
	redisClient *redis.Client
}

func CreateSessionRepository(redisClient *redis.Client) *SessionRepo {
	return &SessionRepo{redisClient: redisClient}
}

func (s *SessionRepo) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	session.Id = uuid.New().String()
	sessionKey := utils.CreateKey(prefix, session.Id)

	j, err := json.Marshal(&session)
	if err != nil {
		return "", errors.WithMessage(err, "Unable to parse session to json")
	}
	if err = s.redisClient.Set(ctx, sessionKey, j, time.Second*time.Duration(expire)).Err(); err != nil {
		return "", errors.Wrap(err, "Unable to set session to db")
	}
	return session.Id, nil
}

func (s *SessionRepo) GetSessionById(ctx context.Context, sessionId string) (*models.Session, error) {
	res, err := s.redisClient.Get(ctx, utils.CreateKey(prefix, sessionId)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get value by key")
	}
	sess := &models.Session{}
	if err = json.Unmarshal(res, &sess); err != nil {
		return nil, errors.Wrap(err, "Unable to unmarshal response")
	}
	return sess, nil
}

func (s *SessionRepo) DeleteById(ctx context.Context, sessionId string) error {
	if err := s.redisClient.Del(ctx, utils.CreateKey(prefix, sessionId)).Err(); err != nil {
		return errors.Wrap(err, "Unable to delete record by id")
	}
	return nil
}

//func (s *SessionRepo) CreateKey(sessionId string) string {
//	return fmt.Sprintf("%s: %s", s.prefix, sessionId)
//}
