package client

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
)

// Account Interface for interacting with the use case repository.
type Account interface {
	CreateAccount(ctx context.Context, account account.Account) (*account.Account, error)
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
}
