package grpc

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dto"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
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

	return dto.ToDomainUser(*data)
}

// CreateUser user registration
func (c *Client) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {

	req := &protobuff.RegisterUserRequest{
		ID:        user.ID().String(),
		Email:     user.Email().String(),
		Password:  user.Password().String(),
		Name:      user.Name().String(),
		Role:      user.Role().String(),
		CreatedAt: user.CreatedAt().Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt().Format(time.RFC3339),
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

	return dto.ToDomainUser(*data)
}

// SetTokenUser install a token User
func (c *Client) SetTokenUser(ctx context.Context, user *user.User) bool {
	return c.authInterceptor.SetToken(user.Token().String())
}

// CountUser get the number of records
func (c *Client) CountUser(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return 0, nil
}
