package account

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, account account.Account) (uuid.UUID, error) {
	return uc.adapterClient.CreateAccount(ctx, account)
}
