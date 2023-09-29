package account

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
)

func (uc *UseCase) Create(ctx context.Context, account account.Account) (*account.Account, error) {
	return uc.adapterClient.CreateAccount(ctx, account)
}
