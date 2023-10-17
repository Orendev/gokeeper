package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Account Interface for interaction between delivery and use case
type Account interface {
	Create(ctx context.Context, account *account.Account) (*account.Account, error)
	Update(ctx context.Context, account *account.Account) (*account.Account, error)
	Delete(ctx context.Context, id uuid.UUID) error
	AccountReader
}

// AccountReader user-readable interface
type AccountReader interface {
	List(ctx context.Context, parameter queryParameter.QueryParameter) (*account.ListAccountViewModel, error)
	Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
