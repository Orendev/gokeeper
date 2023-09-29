package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interaction between delivery and use case

type Account interface {
	Create(accounts ...*account.Account) ([]*account.Account, error)
	Update(ctx context.Context, update account.Account) (*account.Account, error)
	Delete(id uuid.UUID) error

	AccountReader
}

type AccountReader interface {
	List(parameter queryParameter.QueryParameter) ([]*account.Account, error)
	GetByID(id uuid.UUID) (response *account.Account, err error)
}
