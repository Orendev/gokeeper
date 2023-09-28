package account

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(accounts ...*account.Account) ([]*account.Account, error) {
	return uc.adapterStorage.CreateAccount(accounts...)
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

func (uc *UseCase) Delete(id uuid.UUID) error {
	return uc.adapterStorage.DeleteAccount(id)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	return uc.adapterStorage.ListAccount(parameter)
}

func (uc *UseCase) GetByID(ID uuid.UUID) (response *account.Account, err error) {
	return uc.adapterStorage.GetByIDAccount(ID)
}
