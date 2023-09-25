package grpc

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/token"
)

// LoginUser login user and returns the User
func (c *Client) LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
	req := &protobuff.LoginUserRequest{
		Email:    email.String(),
		Password: password.String(),
	}

	res, err := c.KeeperServiceClient.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}

	data, err := dto.FromLoginUserResponseToDto(res)
	if err != nil {
		return nil, err
	}

	data.Password = password.String()

	return toDomainUser(*data)
}

// RegisterUser login user and returns the User
func (c *Client) RegisterUser(ctx context.Context, user user.User) (*user.User, error) {

	req := &protobuff.RegisterUserRequest{
		Email:    user.Email().String(),
		Password: user.Password().String(),
		Name:     user.Name().String(),
		Role:     user.Role().String(),
	}

	res, err := c.KeeperServiceClient.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}

	data, err := dto.FromRegisterUserResponseToDto(res)
	if err != nil {
		return nil, err
	}

	data.Password = user.Password().String()

	return toDomainUser(*data)
}

// SetToken install a token User
func (c Client) SetToken(ctx context.Context, token token.Token) bool {
	return c.authInterceptor.SetToken(token.String())
}
