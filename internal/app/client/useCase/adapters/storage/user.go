package storage

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
)

// User Interface for interacting with the use case repository.
type User interface {
	CreateUser(ctx context.Context, user user.User) (*user.User, error)

	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	FindUser(ctx context.Context, email email.Email) (*user.User, error)
	LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error)
}
