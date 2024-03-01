package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Snorkin/auth_service/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	createUserQuery = `INSERT INTO users (first_name, last_name, email, password, role, avatar) 
		VALUES ($1, $2, $3, $4, $5, COALESCE(NULLIF($6, ''), null)) 
		RETURNING user_id, first_name, last_name, email, password, avatar, created_at, updated_at, role`

	findByEmailQuery = `SELECT * FROM users WHERE email = $1`

	findByIDQuery = `SELECT user_id, email, first_name, last_name, role, avatar, created_at, updated_at FROM users WHERE user_id = $1`
)

type UserRepository struct {
	db *sqlx.DB
}

func CreateUserPgRepo(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	newUser := &models.User{}
	if err := u.db.QueryRowxContext(ctx, createUserQuery, user.Name, user.Surname, user.Email, user.Password).StructScan(newUser); err != nil {
		return nil, errors.Wrap(err, "Unable to create new record in users table")
	}
	return newUser, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	userCheck := &models.User{}
	if err := u.db.GetContext(ctx, userCheck, findByEmailQuery, email); err != nil {
		return nil, errors.Wrap(err, "Unable to find user by email")
	}
	return userCheck, nil
}

func (u *UserRepository) FindById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	userCheck := &models.User{}
	if err := u.db.GetContext(ctx, userCheck, findByIDQuery, id); err != nil {
		return nil, errors.Wrap(err, "Unable to find user by email")
	}
	return userCheck, nil
}
