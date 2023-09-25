package useCase

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interaction between delivery and use case

type Account interface {
	Create(accounts ...*account.Account) ([]*account.Account, error)
	Update(account account.Account) (*account.Account, error)
	Delete(ID uuid.UUID) error

	AccountReader
}

type AccountReader interface {
	List(parameter queryParameter.QueryParameter) ([]*account.Account, error)
	ReadByID(ID uuid.UUID) (response *account.Account, err error)
	Count( /*Тут можно передавать фильтр*/) (uint64, error)
}