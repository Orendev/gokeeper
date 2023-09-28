package storage

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/google/uuid"
)

// Account Interface for interacting with the use case repository.
type Account interface {
	CreateAccount(ctx context.Context, account account.Account) (*account.Account, error)
	UpdateAccount(ctx context.Context, id uuid.UUID, updateFn func(update *account.Account) (*account.Account, error)) (*account.Account, error)
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
	GetAccount(ctx context.Context) (*account.Account, error)
}
