package storage

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Account interface {
	CreateAccount(ctx context.Context, accounts ...*account.Account) ([]*account.Account, error)
	UpdateAccount(ctx context.Context, id uuid.UUID, updateFn func(a *account.Account) (*account.Account, error)) (*account.Account, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error

	AccountReader
}

type AccountReader interface {
	ListAccount(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error)
	GetByIDAccount(ctx context.Context, id uuid.UUID) (response *account.Account, err error)
	CountAccount(ctx context.Context) (uint64, error)
}
