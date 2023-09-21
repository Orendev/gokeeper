package grpc

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/token"
)

// LoginUser login user and returns the access token
func (r *Repository) LoginUser(ctx context.Context, email email.Email, password password.Password) (*token.Token, error) {
	req := &protobuff.LoginUserRequest{
		Email:    email.String(),
		Password: password.String(),
	}

	res, err := r.KeeperServiceClient.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}
	t, err := token.New(res.GetAccessToken())
	if err != nil {
		return nil, err
	}
	return t, nil
}

// FindUser login user and returns the access token
func (r *Repository) FindUser(ctx context.Context, email email.Email) (*user.User, error) {
	panic("implement me")
}

// CreateUser login user and returns the access token
func (r *Repository) CreateUser(ctx context.Context, userObject user.User) (*user.User, error) {

	req := &protobuff.CreateUserRequest{
		Email:    userObject.Email().String(),
		Password: userObject.Password().String(),
	}

	res, err := r.KeeperServiceClient.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return r.toDomainUser(res, userObject)
}
