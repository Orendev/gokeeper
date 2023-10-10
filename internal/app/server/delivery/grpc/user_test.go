package grpc

import (
	"errors"
	"testing"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	useCaseUser "github.com/Orendev/gokeeper/internal/pkg/useCase/user"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/role"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

var (
	dataUser     = make(map[string]*user.User)
	registerUser *user.User

	registerReqUser *protobuff.RegisterUserRequest
	registerResUser *protobuff.RegisterUserResponse

	loginReqUser *protobuff.LoginUserRequest
	loginResUser *protobuff.LoginUserResponse
)

func testMainUser() {
	passwordUser, _ := password.New("supersecret")
	emailUser, _ := email.New("admin@ya.ru")
	nameUser, _ := name.New("Ivanov")
	roleUser, _ := role.New("user")

	registerUser, _ = user.NewWithID(
		uuid.New(),
		*passwordUser,
		*emailUser,
		*nameUser,
		*roleUser,
		time.Now().UTC(),
		time.Now().UTC(),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0

	registerReqUser = &protobuff.RegisterUserRequest{
		ID:        registerUser.ID().String(),
		Email:     registerUser.Email().String(),
		Password:  registerUser.Password().String(),
		Name:      registerUser.Name().String(),
		Role:      registerUser.Role().String(),
		CreatedAt: registerUser.CreatedAt().Format(time.RFC3339),
		UpdatedAt: registerUser.UpdatedAt().Format(time.RFC3339),
	}

	accessToken, _ := jwtManager.Generate(registerUser.ID())

	registerResUser = &protobuff.RegisterUserResponse{
		ID:          registerUser.ID().String(),
		Email:       registerUser.Email().String(),
		AccessToken: accessToken,
		Name:        registerUser.Name().String(),
		Role:        registerUser.Role().String(),
		CreatedAt:   registerUser.CreatedAt().String(),
		UpdatedAt:   registerUser.UpdatedAt().String(),
	}

	loginReqUser = &protobuff.LoginUserRequest{Email: emailUser.String(), Password: passwordUser.String()}

	loginResUser = &protobuff.LoginUserResponse{
		ID:          registerUser.ID().String(),
		Email:       registerUser.Email().String(),
		AccessToken: accessToken,
		Name:        registerUser.Name().String(),
		Role:        registerUser.Role().String(),
		CreatedAt:   registerUser.CreatedAt().String(),
		UpdatedAt:   registerUser.UpdatedAt().String(),
	}

}

func initTestUseCaseUser(t *testing.T) {
	assertion := assert.New(t)
	storageRepository.On("CreateUser",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, user *user.User) (*user.User, error) {
			assertion.Equal(user.ID(), registerUser.ID())
			dataUser[user.Email().String()] = user

			return user, nil
		}, func(ctx context.Context, user *user.User) (*user.User, error) {
			return nil, errors.New("user creation error")
		})

	storageRepository.On("LoginUser",
		mock.Anything,
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
			u, ok := dataUser[email.String()]
			if !ok {
				return nil, errors.New("incorrect email/password")
			}

			return u, nil
		}, func(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
			return nil, errors.New("incorrect email/password")
		})

	storageRepository.On("SetTokenUser",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, user *user.User) bool {
			return true
		}, func(ctx context.Context, user *user.User) bool {
			return false
		})

}

func TestDeliveryUser(t *testing.T) {
	initTestUseCaseUser(t)
	ucUser = useCaseUser.New(storageRepository, useCaseUser.Options{})
	option := Options{}
	deliveryGRPC := New(ucUser, ucAccount, ucCard, ucText, ucBinary, jwtManager, option)
	assertion := assert.New(t)

	t.Run("positive register user", func(t *testing.T) {
		result, err := deliveryGRPC.RegisterUser(context.Background(), registerReqUser)
		assertion.NoError(err)
		assertion.Equal(result.AccessToken, registerResUser.AccessToken)
	})

	t.Run("positive login user", func(t *testing.T) {
		result, err := deliveryGRPC.LoginUser(context.Background(), loginReqUser)
		assertion.NoError(err)
		assertion.Equal(result.AccessToken, loginResUser.AccessToken)
	})

	t.Run("negative login user", func(t *testing.T) {
		result, err := deliveryGRPC.LoginUser(context.Background(), &protobuff.LoginUserRequest{Email: "t@y.ru", Password: "secret"})
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

}
