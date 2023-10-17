package account

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, account *account.Account) (*account.Account, error) {
	return uc.adapterStorage.CreateAccount(ctx, account)
}

func (uc *UseCase) Update(ctx context.Context, account *account.Account) (*account.Account, error) {
	return uc.adapterStorage.UpdateAccount(ctx, account)
}

func (uc *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.adapterStorage.DeleteAccount(ctx, id)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) (*account.ListAccountViewModel, error) {
	return uc.adapterStorage.ListAccount(ctx, parameter)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountAccount(ctx, parameter)
}
