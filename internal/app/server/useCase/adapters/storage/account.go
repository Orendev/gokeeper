package storage

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Account interface {
	CreateAccount(accounts ...*account.Account) ([]*account.Account, error)
	UpdateAccount(ctx context.Context, id uuid.UUID, updateFn func(a *account.Account) (*account.Account, error)) (*account.Account, error)
	DeleteAccount(id uuid.UUID) error

	AccountReader
}

type AccountReader interface {
	ListAccount(parameter queryParameter.QueryParameter) ([]*account.Account, error)
	GetByIDAccount(id uuid.UUID) (response *account.Account, err error)
}
