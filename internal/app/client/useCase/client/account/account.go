package account

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
)

func (uc *UseCase) CreateAccount(ctx context.Context, account account.Account) (*account.Account, error) {
	return uc.adapterClient.CreateAccount(ctx, account)
}
