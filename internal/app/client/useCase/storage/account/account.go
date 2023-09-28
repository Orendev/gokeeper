package account

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"time"
)

// Create creating account
func (uc *UseCase) Create(ctx context.Context, account account.Account) (*account.Account, error) {
	return uc.adapterStorage.CreateAccount(ctx, account)
}

// Update update account
func (uc *UseCase) Update(ctx context.Context, update account.Account) (*account.Account, error) {
	return uc.adapterStorage.UpdateAccount(ctx, update.ID(), func(old *account.Account) (*account.Account, error) {
		return account.NewWithID(old.ID(), update.Title(), update.Login(), update.Password(), update.URL(), update.Comment(), update.IsDeleted(), old.CreatedAt(), time.Now().UTC())
	})
}

// Get let's get the account
func (uc *UseCase) Get(ctx context.Context) (*account.Account, error) {
	return uc.adapterStorage.GetAccount(ctx)
}
