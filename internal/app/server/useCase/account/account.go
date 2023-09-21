package account

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(accounts ...*account.Account) ([]*account.Account, error) {
	return uc.adapterStorage.CreateAccount(accounts...)
}

func (uc *UseCase) Update(accountUpdate account.Account) (*account.Account, error) {
	return uc.adapterStorage.UpdateAccount(accountUpdate.ID(), func(oldAccount *account.Account) (*account.Account, error) {
		return account.NewWithID(
			oldAccount.ID(),
			oldAccount.UserID(),
			accountUpdate.Title(),
			accountUpdate.Login(),
			accountUpdate.Password(),
			accountUpdate.WebAddress(),
			accountUpdate.Comment(),
			accountUpdate.Version(),
			oldAccount.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteAccount(ID)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	return uc.adapterStorage.ListAccount(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *account.Account, err error) {
	return uc.adapterStorage.ReadAccountByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountAccount()
}
