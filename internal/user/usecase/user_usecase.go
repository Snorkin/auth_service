package usecase

import (
	"context"
	"github.com/Snorkin/auth_service/internal/models"
	"github.com/Snorkin/auth_service/internal/user"
	"github.com/Snorkin/auth_service/pkg/grpc_errors"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userPgRepo user.PgRepository
}

func CreateUserUseCase(userRepo user.PgRepository) user.Useсase {
	return &UserUsecase{userPgRepo: userRepo}
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

func (u *UserUsecase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	userCheck, err := u.userPgRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, grpc_errors.ErrNotFound
	}
	userCheck.ClearPassword()
	return userCheck, nil
}

func (u *UserUsecase) FindById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
