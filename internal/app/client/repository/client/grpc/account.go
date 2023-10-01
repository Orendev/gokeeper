package grpc

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (c *Client) CreateAccount(ctx context.Context, account account.Account) (uuid.UUID, error) {
	ac := &protobuff.Account{
		ID:       account.ID().String(),
		Title:    account.Title().String(),
		Login:    account.Login().String(),
		Password: account.Password().String(),
		URL:      account.URL().String(),
		Comment:  account.Comment().String(),
	}
	req := &protobuff.CreateAccountRequest{Data: ac}

	res, err := c.KeeperServiceClient.CreateAccount(ctx, req)
	if err != nil {
		return uuid.Nil, err
	}

	data, err := dto.FromCreateAccountResponseToDto(res)
	if err != nil {
		return uuid.Nil, err
	}

	return toAccountId(data.ID)
}

func (c Client) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	req := &protobuff.DeleteAccountRequest{
		ID: id.String(),
	}

	res, err := c.KeeperServiceClient.DeleteAccount(ctx, req)
	if err != nil {
		return err
	}

	_, err = dto.FromDeleteAccountResponseToDto(res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateAccount(ctx context.Context, account account.Account) (uuid.UUID, error) {
	ac := &protobuff.Account{
		ID:       account.ID().String(),
		Title:    account.Title().String(),
		Login:    account.Login().String(),
		Password: account.Password().String(),
		URL:      account.URL().String(),
		Comment:  account.Comment().String(),
	}
	req := &protobuff.UpdateAccountRequest{Data: ac}

	res, err := c.KeeperServiceClient.UpdateAccount(ctx, req)
	if err != nil {
		return uuid.Nil, err
	}

	data, err := dto.FromUpdateAccountResponseToDto(res)
	if err != nil {
		return uuid.Nil, err
	}

	return toAccountId(data.ID)
}

func (c *Client) ListAccount(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error) {

	req := &protobuff.ListAccountRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	res, err := c.KeeperServiceClient.ListAccount(ctx, req)

	if err != nil {
		return nil, err
	}

	listAccount, err := dto.FromListAccountResponseToDto(res)
	if err != nil {
		return nil, err
	}

	return toDomainAccounts(*listAccount)
}
