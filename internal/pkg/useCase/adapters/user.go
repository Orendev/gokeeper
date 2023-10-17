package adapters

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

// User Interface for interacting with the use case repository.
type User interface {
	CreateUser(ctx context.Context, user *user.User) (*user.User, error)
	SetTokenUser(ctx context.Context, user *user.User) bool
	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error)
	CountUser(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
