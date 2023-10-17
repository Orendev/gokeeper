package user

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

// Create creating a user
func (uc *UseCase) Create(ctx context.Context, user *user.User) (*user.User, error) {
	return uc.adapterStorage.CreateUser(ctx, user)
}

// SetToken install the token
func (uc UseCase) SetToken(ctx context.Context, user *user.User) bool {
	return uc.adapterStorage.SetTokenUser(ctx, user)
}

// Login user authorization
func (uc *UseCase) Login(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
	return uc.adapterStorage.LoginUser(ctx, email, password)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountUser(ctx, parameter)
}
