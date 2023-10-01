package account

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, accounts ...*account.Account) ([]*account.Account, error) {
	return uc.adapterStorage.CreateAccount(ctx, accounts...)
}

func (uc *UseCase) Update(ctx context.Context, update account.Account) (*account.Account, error) {
	return uc.adapterStorage.UpdateAccount(ctx, update.ID(), func(oldAccount *account.Account) (*account.Account, error) {
		return account.NewWithID(
			oldAccount.ID(),
			oldAccount.UserID(),
			update.Title(),
			update.Login(),
			update.Password(),
			update.URL(),
			update.Comment(),
			oldAccount.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.adapterStorage.DeleteAccount(ctx, id)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	return uc.adapterStorage.ListAccount(ctx, parameter)
}

func (uc *UseCase) GetByID(ctx context.Context, ID uuid.UUID) (response *account.Account, err error) {
	return uc.adapterStorage.GetByIDAccount(ctx, ID)
}

func (uc *UseCase) Count(ctx context.Context) (uint64, error) {
	return uc.adapterStorage.CountAccount(ctx)
}
