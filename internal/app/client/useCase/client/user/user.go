package user

import (
	"context"

	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
)

func (uc *UseCase) Login(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {

	return uc.adapterClient.LoginUser(ctx, email, password)
}

func (uc *UseCase) Register(ctx context.Context, user user.User) (*user.User, error) {

	return uc.adapterClient.RegisterUser(ctx, user)
}

func (uc *UseCase) SetToken(user user.User) bool {
	return uc.adapterClient.SetToken(user.Token())
}
