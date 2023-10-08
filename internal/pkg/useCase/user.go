package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

// User Interface for interaction between delivery and use case
type User interface {
	Create(ctx context.Context, user *user.User) (*user.User, error)
	SetToken(ctx context.Context, user *user.User) bool
	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	Login(ctx context.Context, email email.Email, password password.Password) (*user.User, error)
	Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
