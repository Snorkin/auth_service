package grpc_errors

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrNotFound         = errors.New("Not found")
	ErrNotCtxMetaData   = errors.New("No context metadata")
	ErrInvalidSessionId = errors.New("Invalid session Id")
	ErrEmailExists      = errors.New("Email already exists")
	ErrPermissionDenied = errors.New("Permission denied")
)

func GetGRPCStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, redis.Nil):
		return codes.NotFound
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case errors.Is(err, ErrEmailExists):
		return codes.AlreadyExists
	case errors.Is(err, ErrNotCtxMetaData):
		return codes.Unauthenticated
	case errors.Is(err, ErrInvalidSessionId):
		return codes.PermissionDenied
	case errors.Is(err, ErrNotFound):
		return codes.NotFound
	case errors.Is(err, ErrPermissionDenied):
		return codes.PermissionDenied
	}
	return codes.Internal
}
