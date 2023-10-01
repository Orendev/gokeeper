package storage

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Account Interface for interaction between delivery and use case
type Account interface {
	Create(ctx context.Context, account account.Account) (*account.Account, error)
	Update(ctx context.Context, update account.Account) (*account.Account, error)
	Delete(ctx context.Context, id uuid.UUID) error
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
	GetByID(ctx context.Context, id uuid.UUID) (*account.Account, error)
	List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error)
}
