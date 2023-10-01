package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interaction between delivery and use case

type Account interface {
	Create(ctx context.Context, accounts ...*account.Account) ([]*account.Account, error)
	Update(ctx context.Context, update account.Account) (*account.Account, error)
	Delete(ctx context.Context, id uuid.UUID) error

	AccountReader
}

type AccountReader interface {
	List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error)
	GetByID(ctx context.Context, id uuid.UUID) (response *account.Account, err error)
	Count(ctx context.Context) (uint64, error)
}
