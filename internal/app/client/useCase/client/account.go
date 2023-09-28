package client

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
)

type Account interface {
	Create(ctx context.Context, account account.Account) (*account.Account, error)
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
}
