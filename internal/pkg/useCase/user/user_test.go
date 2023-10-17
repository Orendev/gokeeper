package user

import (
	"context"
	"os"
	"testing"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/mocks"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	storageRepository = new(mocks.Storage)
	ucUser            *UseCase
	data              = make(map[string]*user.User)
	createUser        *user.User
	updateUser        *user.User
	parameter         = queryParameter.QueryParameter{}
)

func TestMain(m *testing.M) {
	passwordObj, _ := password.New("test password")
	emailObj, _ := email.New("test@ya.ru")
	nameObj, _ := name.New("test")
	tokenObj := token.New("test")

	createUser, _ = user.New(
		*passwordObj,
		*emailObj,
		*nameObj,
		*tokenObj,
	)

	tokenUpdateObj := token.New("test update")

	updateUser, _ = user.New(
		*passwordObj,
		*emailObj,
		*nameObj,
		*tokenUpdateObj,
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0

	os.Exit(m.Run())
}

func initTestUseCaseAccount(t *testing.T) {
	assertion := assert.New(t)

	storageRepository.On("CreateUser",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, user *user.User) *user.User {
			assertion.Equal(user, createUser)
			data[user.Email().String()] = user

			return user
		}, func(ctx context.Context, account *user.User) error {
			return nil
		})

	storageRepository.On("SetTokenUser",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, user *user.User) bool {
			assertion.Equal(user, updateUser)

			_, ok := data[user.Email().String()]
			if !ok {
				return false
			}
			data[user.Email().String()] = user

			return true
		})

	storageRepository.On("LoginUser",
		mock.Anything,
		mock.AnythingOfType("email.Email"),
		mock.AnythingOfType("password.Password")).
		Return(func(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
			u, ok := data[email.String()]
			if !ok {
				return nil, repository.ErrNotFoundUser
			}

			return u, nil
		}, func(ctx context.Context, email email.Email, password password.Password) error {
			return nil
		})

	storageRepository.On("CountUser",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {

			return uint64(len(data)), nil

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestAccount(t *testing.T) {
	initTestUseCaseAccount(t)
	ucUser = New(storageRepository, Options{})
	assertion := assert.New(t)

	t.Run("create account", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucUser.Create(ctx, createUser)
		assertion.NoError(err)
		assertion.Equal(result, createUser)
	})

	t.Run("set token user", func(t *testing.T) {
		var ctx = context.Background()

		result := ucUser.SetToken(ctx, updateUser)

		assertion.True(result)
	})

	t.Run("login user", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucUser.Login(ctx, createUser.Email(), createUser.Password())
		assertion.NoError(err)
		assertion.Equal(result.Email(), createUser.Email())

	})

	t.Run("count user", func(t *testing.T) {
		var ctx = context.Background()
		total := uint64(len(data))
		result, err := ucUser.Count(ctx, parameter)
		assertion.NoError(err)
		assertion.Equal(result, total)
	})
}
