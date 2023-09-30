package storage

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
)

// Account Interface for interaction between delivery and use case
type Account interface {
	Create(ctx context.Context, account account.Account) (*account.Account, error)
	Update(ctx context.Context, update account.Account) (*account.Account, error)
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
	Get(ctx context.Context) (*account.Account, error)
}
