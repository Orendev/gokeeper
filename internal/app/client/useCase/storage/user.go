package storage

import (
	"context"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/google/uuid"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
)

// User Interface for interaction between delivery and use case
type User interface {
	Add(ctx context.Context, user user.User) (*user.User, error)
	UpdateToken(ctx context.Context, id uuid.UUID, token token.Token) (*user.User, error)
	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	Get(ctx context.Context) (*user.User, error)
}
