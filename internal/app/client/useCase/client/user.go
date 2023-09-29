package client

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
)

// User Interface for interaction between delivery and use case
type User interface {
	Register(ctx context.Context, user user.User) (*user.User, error)
	SetToken(user user.User) bool
	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	Login(ctx context.Context, email email.Email, password password.Password) (*user.User, error)
}
