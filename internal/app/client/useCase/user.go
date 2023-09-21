package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
)

// User Interface for interaction between delivery and use case
type User interface {
	Create(ctx context.Context, user user.User) (*user.User, error)

	UserReader
}

type UserReader interface {
	Login(ctx context.Context, email email.Email, password password.Password) (*user.User, error)
	Find(ctx context.Context, email email.Email) (*user.User, error)
}
