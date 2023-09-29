package client

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/token"
)

// User Interface for interacting with the use case repository.
type User interface {
	RegisterUser(ctx context.Context, user user.User) (*user.User, error)
	SetToken(token token.Token) bool
	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error)
}
