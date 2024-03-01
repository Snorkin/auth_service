package usecase

import (
	"context"
	"github.com/Snorkin/auth_service/internal/models"
	"github.com/Snorkin/auth_service/internal/user"
	"github.com/Snorkin/auth_service/pkg/grpc_errors"
	"github.com/Snorkin/auth_service/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	userCacheDuration = 3600
)

type UserUsecase struct {
	logger        logger.Logger
	userPgRepo    user.PgRepository
	userRedisRepo user.RedisRepository
}

func CreateUserUseCase(logger logger.Logger, userRepo user.PgRepository, userRedisRepo user.RedisRepository) *UserUsecase {
	return &UserUsecase{userPgRepo: userRepo, userRedisRepo: userRedisRepo, logger: logger}
}

func (u *UserUsecase) Register(ctx context.Context, user *models.User) (*models.User, error) {
	checkUser, err := u.userPgRepo.FindByEmail(ctx, user.Email)
	if checkUser != nil || err == nil {
		return nil, grpc_errors.ErrEmailExists
	}
	if checkUser != nil || err != nil {
		return nil, err //grpc err email already exists here
	}
	return u.userPgRepo.CreateUser(ctx, user)
}

func (u *UserUsecase) Login(ctx context.Context, email string, password string) (*models.User, error) {
	checkUser, err := u.userPgRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, grpc_errors.ErrNotFound
	}
	if err := checkUser.ComparePasswords(password); err != nil {
		return nil, grpc_errors.ErrPermissionDenied
	}
	return checkUser, nil
}

// FindByEmail TODO: implement cache
func (u *UserUsecase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	userCheck, err := u.userPgRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, grpc_errors.ErrNotFound
	}
	userCheck.ClearPassword()
	return userCheck, nil
}

func (u *UserUsecase) FindById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	cachedUser, err := u.userRedisRepo.GetByIdCtx(ctx, id.String())
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, grpc_errors.ErrNotFound
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	userCheck, err := u.userPgRepo.FindById(ctx, id)
	if err != nil {
		return nil, grpc_errors.ErrNotFound
	}

	if err := u.userRedisRepo.SetUserCtx(ctx, userCheck.UserID.String(), userCacheDuration, userCheck); err != nil {
		u.logger.Errorf("Unable to set user to cache db")
	}

	return userCheck, nil
}
