package repository

import (
	"context"
	"encoding/json"
	"github.com/Snorkin/auth_service/internal/models"
	"github.com/Snorkin/auth_service/pkg/grpc_errors"
	"github.com/Snorkin/auth_service/pkg/utils"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	prefix = "user:"
)

type UserRedisRepo struct {
	redisClient *redis.Client
}

func CreateUserRedisRepository(redisClient *redis.Client) *UserRedisRepo {
	return &UserRedisRepo{
		redisClient: redisClient,
	}
}

func (u *UserRedisRepo) GetByIdCtx(ctx context.Context, key string) (*models.User, error) {
	userBytes, err := u.redisClient.Get(ctx, utils.CreateKey(prefix, key)).Bytes()
	if err != nil {
		if err != redis.Nil {
			return nil, grpc_errors.ErrNotFound
		}
		return nil, err
	}
	checkUser := &models.User{}
	if err = json.Unmarshal(userBytes, checkUser); err != nil {
		return nil, err
	}
	return checkUser, nil
}

func (u *UserRedisRepo) SetUserCtx(ctx context.Context, key string, seconds int, user *models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return u.redisClient.Set(ctx, utils.CreateKey(prefix, key), userBytes, time.Second*time.Duration(seconds)).Err()
}

func (u *UserRedisRepo) DeleteUserCtx(ctx context.Context, key string) error {
	return u.redisClient.Del(ctx, utils.CreateKey(prefix, key)).Err()
}
