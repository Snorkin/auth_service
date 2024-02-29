package usecase

import (
	"context"
	"github.com/Snorkin/auth_service/internal/models"
	"github.com/Snorkin/auth_service/internal/user"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userPgRepo user.PgRepository
}

func CreateUserUseCase(userRepo user.PgRepository) *UserUsecase {
	return &UserUsecase{userPgRepo: userRepo}
}

func (u *UserUsecase) Register(ctx context.Context, user *models.User) (*models.User, error) {
	checkUser, err := u.userPgRepo.FindByEmail(ctx, user)
	if checkUser != nil || err != nil {
		return nil, err //grpc err email already exists here
	}
	return u.userPgRepo.Create(ctx, user)
}

func (u *UserUsecase) Login(ctx context.Context, email string, password string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUsecase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUsecase) FindById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
