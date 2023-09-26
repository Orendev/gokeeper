package storage

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
)

// User Interface for interaction between delivery and use case
type User interface {
	Add(ctx context.Context, user user.User) (*user.User, error)
	UpdateToken(ctx context.Context, update user.User) (*user.User, error)
	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	Get(ctx context.Context) (*user.User, error)
}
