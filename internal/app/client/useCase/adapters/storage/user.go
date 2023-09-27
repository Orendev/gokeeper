package storage

import (
	"context"
	"github.com/google/uuid"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
)

// User Interface for interacting with the use case repository.
type User interface {
	AddUser(ctx context.Context, user user.User) (*user.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, updateFn func(u *user.User) (*user.User, error)) (*user.User, error)
	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	GetUser(ctx context.Context) (*user.User, error)
}
