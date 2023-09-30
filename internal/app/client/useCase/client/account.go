package client

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/google/uuid"
)

// Account Interface for interaction between delivery and use case
type Account interface {
	Create(ctx context.Context, account account.Account) (uuid.UUID, error)
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
}
