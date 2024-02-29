package user

import (
	"context"

	"github.com/Snorkin/auth_service/internal/models"
	"github.com/google/uuid"
)

type PgRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, user *models.User) (*models.User, error)
	FindById(ctx context.Context, user *models.User) (*models.User, error)
}

type Useсase interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*models.User, error)
}
