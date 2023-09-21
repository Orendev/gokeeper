package user

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
)

func (uc *UseCase) Create(ctx context.Context, user user.User) (*user.User, error) {
	return uc.adapterStorage.CreateUser(ctx, user)
}

func (uc *UseCase) Find(ctx context.Context, email email.Email) (*user.User, error) {
	return uc.adapterStorage.FindUser(ctx, email)
}

func (uc *UseCase) Login(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
	return uc.adapterStorage.LoginUser(ctx, email, password)
}
