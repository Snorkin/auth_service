package repository

import (
	"context"

	"github.com/Snorkin/auth_service/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserPgRepo(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

func (u *UserRepository) FindById(ctx context.Context, user *models.User) (*models.User, error) {
	return &models.User{}, nil
}
