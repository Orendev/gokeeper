package client

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Account Interface for interacting with the use case repository.
type Account interface {
	CreateAccount(ctx context.Context, account account.Account) (uuid.UUID, error)
	UpdateAccount(ctx context.Context, account account.Account) (uuid.UUID, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
	ListAccount(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error)
}
