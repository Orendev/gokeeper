package grpc

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/google/uuid"
)

func (c *Client) CreateAccount(ctx context.Context, account account.Account) (uuid.UUID, error) {
	req := &protobuff.CreateAccountRequest{
		ID:       account.ID().String(),
		Title:    account.Title().String(),
		Login:    account.Login().String(),
		Password: account.Password().String(),
		URL:      account.URL().String(),
		Comment:  account.Comment().String(),
	}

	res, err := c.KeeperServiceClient.CreateAccount(ctx, req)
	if err != nil {
		return uuid.Nil, err
	}

	data, err := dto.FromCreateAccountResponseToDto(res)
	if err != nil {
		return uuid.Nil, err
	}

	return toAccountId(*data)
}
