package grpc

import (
	"fmt"
	"testing"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	useCaseAccount "github.com/Orendev/gokeeper/internal/pkg/useCase/account"

	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

var (
	accounts             []*account.Account
	dataAccount          = make(map[uuid.UUID]*account.Account)
	listAccountViewModel = &account.ListAccountViewModel{}
	createAccount        *account.Account
	updateAccount        *account.Account
	totalAccount         uint64
	createReqAccount     *protobuff.CreateAccountRequest
	createResAccount     *protobuff.CreateAccountResponse

	updateReqAccount *protobuff.UpdateAccountRequest
	updateResAccount *protobuff.UpdateAccountResponse

	deleteReqAccount *protobuff.DeleteAccountRequest
	deleteResAccount *protobuff.DeleteAccountResponse

	listReqAccount *protobuff.ListAccountRequest
	listResAccount *protobuff.ListAccountResponse
)

func testMainAccount() {
	titleObj, _ := title.New("Test title")
	createAccount, _ = account.New(
		userID,
		*titleObj,
		[]byte("login"),
		[]byte("password"),
		[]byte("ya.ru"),
		[]byte("comment"),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0

	createAccountProto := &protobuff.Account{
		ID:        createAccount.ID().String(),
		Title:     createAccount.Title().String(),
		Login:     createAccount.Login(),
		Password:  createAccount.Password(),
		URL:       createAccount.URL(),
		Comment:   createAccount.Comment(),
		CreatedAt: createAccount.CreatedAt().Format(time.RFC3339),
		UpdatedAt: createAccount.UpdatedAt().Format(time.RFC3339),
		UserID:    createAccount.UserID().String(),
	}
	createReqAccount = &protobuff.CreateAccountRequest{Data: createAccountProto}

	createResAccount = &protobuff.CreateAccountResponse{
		ID: createAccount.ID().String(),
	}

	titleUpdateObj, _ := title.New("Test Update")

	updateAccount, _ = account.NewWithID(
		createAccount.ID(),
		userID,
		*titleUpdateObj,
		[]byte("login update"),
		[]byte("password update"),
		[]byte("practicum.ru"),
		[]byte("comment update"),
		createAccount.CreatedAt().UTC(),
		time.Now().UTC(),
	)

	updateAccountProto := &protobuff.Account{
		ID:        updateAccount.ID().String(),
		Title:     updateAccount.Title().String(),
		Login:     updateAccount.Login(),
		Password:  updateAccount.Password(),
		URL:       updateAccount.URL(),
		Comment:   updateAccount.Comment(),
		CreatedAt: updateAccount.CreatedAt().Format(time.RFC3339),
		UpdatedAt: updateAccount.UpdatedAt().Format(time.RFC3339),
		UserID:    updateAccount.UserID().String(),
	}

	updateReqAccount = &protobuff.UpdateAccountRequest{Data: updateAccountProto}

	updateResAccount = &protobuff.UpdateAccountResponse{
		ID: updateAccount.ID().String(),
	}

	deleteReqAccount = &protobuff.DeleteAccountRequest{
		ID: updateAccount.ID().String(),
	}
	deleteResAccount = &protobuff.DeleteAccountResponse{
		ID: updateAccount.ID().String(),
	}

	listReqAccount = &protobuff.ListAccountRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	totalAccount = 1

	listAccountProto := []*protobuff.Account{}

	listAccountProto = append(listAccountProto, updateAccountProto)
	listResAccount = &protobuff.ListAccountResponse{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Total:  totalAccount,
		Data:   listAccountProto,
	}
}

func initTestUseCaseAccount(t *testing.T) {
	assertion := assert.New(t)
	storageRepository.On("CreateAccount",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, account *account.Account) *account.Account {
			assertion.Equal(account.ID(), updateAccount.ID())
			accounts = append(accounts, account)
			listAccountViewModel.Data = accounts
			dataAccount[account.ID()] = account

			return account
		}, func(ctx context.Context, account *account.Account) error {
			return nil
		})

	storageRepository.On("UpdateAccount",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, account *account.Account) *account.Account {
			assertion.Equal(account.ID(), updateAccount.ID())
			if len(accounts) > 0 {
				accounts[0] = account
			} else {
				accounts = append(accounts, account)
			}

			dataAccount[account.ID()] = account
			listAccountViewModel.Data = accounts

			return account
		}, func(ctx context.Context, account *account.Account) error {
			return nil
		})

	storageRepository.On("DeleteAccount",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, id uuid.UUID) error {
			if _, ok := dataAccount[id]; !ok {
				return repository.ErrDataNotFound
			}
			return nil
		}, func(ctx context.Context, id uuid.UUID) error {
			return nil
		})

	storageRepository.On("ListAccount",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) *account.ListAccountViewModel {

			listAccountViewModel.Limit = parameter.Pagination.Limit
			listAccountViewModel.Offset = parameter.Pagination.Offset
			listAccountViewModel.Total = uint64(len(listAccountViewModel.Data))

			return listAccountViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestDeliveryAccount(t *testing.T) {
	initTestUseCaseAccount(t)
	ucAccount = useCaseAccount.New(storageRepository, useCaseAccount.Options{})
	option := Options{}
	deliveryGRPC := New(ucUser, ucAccount, ucCard, ucText, ucBinary, jwtManager, option)
	assertion := assert.New(t)

	t.Run("positive create account", func(t *testing.T) {
		result, err := deliveryGRPC.CreateAccount(ctx, createReqAccount)
		assertion.NoError(err)
		assertion.Equal(result, createResAccount)
	})

	t.Run("negative create account", func(t *testing.T) {
		result, err := deliveryGRPC.CreateAccount(context.Background(), createReqAccount)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive update account", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateAccount(ctx, updateReqAccount)
		assertion.NoError(err)
		assertion.Equal(result, updateResAccount)
	})

	t.Run("negative update account", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateAccount(context.Background(), updateReqAccount)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive delete account", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteAccount(ctx, deleteReqAccount)
		assertion.NoError(err)
		assertion.Equal(result, deleteResAccount)
	})

	t.Run("negative delete account", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteAccount(context.Background(), &protobuff.DeleteAccountRequest{ID: uuid.New().String()})
		assertion.Errorf(err, "")
		assertion.Nil(result)
	})

	t.Run("positive list account", func(t *testing.T) {
		totalAccount = uint64(len(listAccountViewModel.Data))
		result, err := deliveryGRPC.ListAccount(ctx, listReqAccount)
		assertion.NoError(err)
		fmt.Println("tot", totalAccount)
		assertion.Equal(result.Data[0], listResAccount.Data[0])
		assertion.Equal(result.Total, totalAccount)
	})

	t.Run("negative list account", func(t *testing.T) {
		result, err := deliveryGRPC.ListAccount(context.Background(), listReqAccount)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})
}
