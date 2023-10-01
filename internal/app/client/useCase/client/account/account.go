package account

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, account account.Account) (uuid.UUID, error) {
	return uc.adapterClient.CreateAccount(ctx, account)
}

func (uc *UseCase) Update(ctx context.Context, account account.Account) (uuid.UUID, error) {
	return uc.adapterClient.UpdateAccount(ctx, account)
}

func (uc *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.adapterClient.DeleteAccount(ctx, id)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	return uc.adapterClient.ListAccount(ctx, parameter)
}
