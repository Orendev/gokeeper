package account

import (
	"context"
	"os"
	"testing"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/mocks"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	storageRepository    = new(mocks.Storage)
	ucAccount            *UseCase
	data                 = make(map[uuid.UUID]*account.Account)
	accounts             []*account.Account
	listAccountViewModel = &account.ListAccountViewModel{}
	createAccount        *account.Account
	updateAccount        *account.Account
	parameter            = queryParameter.QueryParameter{}
	total                uint64
)

func TestMain(m *testing.M) {
	titleObj, _ := title.New("Test title")
	userID := uuid.New()
	login := []byte("test login")
	password := []byte("test pasword")
	url := []byte("ya.ru")
	comment := []byte("test comment")

	createAccount, _ = account.New(
		userID,
		*titleObj,
		login,
		password,
		url,
		comment,
	)

	titleUpdateObj, _ := title.New("Test Update")

	updateAccount, _ = account.New(
		userID,
		*titleUpdateObj,
		login,
		password,
		url,
		comment,
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0
	total = 1

	os.Exit(m.Run())
}

func initTestUseCaseAccount(t *testing.T) {
	assertion := assert.New(t)

	storageRepository.On("CreateAccount",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, account *account.Account) *account.Account {
			assertion.Equal(account, createAccount)
			accounts = append(accounts, account)
			listAccountViewModel.Data = accounts
			data[account.ID()] = account

			return account
		}, func(ctx context.Context, account *account.Account) error {
			return nil
		})

	storageRepository.On("UpdateAccount",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, account *account.Account) *account.Account {
			assertion.Equal(account, updateAccount)
			data[account.ID()] = account
			listAccountViewModel.Data = accounts

			return account
		}, func(ctx context.Context, account *account.Account) error {
			return nil
		})

	storageRepository.On("DeleteAccount",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, id uuid.UUID) error {
			if _, ok := data[id]; !ok {
				return repository.ErrAccountNotFound
			}
			return nil
		}, func(ctx context.Context, id uuid.UUID) error {
			return nil
		})

	storageRepository.On("ListAccount",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) *account.ListAccountViewModel {
			t := len(listAccountViewModel.Data)
			listAccountViewModel.Limit = parameter.Pagination.Limit
			listAccountViewModel.Offset = parameter.Pagination.Offset
			listAccountViewModel.Total = uint64(t)

			return listAccountViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestAccount(t *testing.T) {
	initTestUseCaseAccount(t)
	ucAccount = New(storageRepository, Options{})
	assertion := assert.New(t)

	t.Run("create account", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucAccount.Create(ctx, createAccount)
		assertion.NoError(err)
		assertion.Equal(result, createAccount)
	})

	t.Run("update account", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucAccount.Update(ctx, updateAccount)
		assertion.NoError(err)
		assertion.Equal(result, updateAccount)
	})

	t.Run("delete account", func(t *testing.T) {
		var ctx = context.Background()

		err := ucAccount.Delete(ctx, createAccount.ID())
		assertion.NoError(err)

	})

	t.Run("list account", func(t *testing.T) {
		var ctx = context.Background()
		total = uint64(len(listAccountViewModel.Data))
		result, err := ucAccount.List(ctx, parameter)
		assertion.NoError(err)
		assertion.Equal(result.Data, listAccountViewModel.Data)
		assertion.Equal(result.Total, total)
	})
}
