package account

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
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

// Delete the account
func (uc *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.adapterStorage.DeleteAccount(ctx, id)
}

// GetByID let's get the account
func (uc *UseCase) GetByID(ctx context.Context, id uuid.UUID) (*account.Account, error) {
	return uc.adapterStorage.GetByIDAccount(ctx, id)
}

// List creating account
func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	return uc.adapterStorage.ListAccount(ctx, parameter)
}
