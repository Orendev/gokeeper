package grpc

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dto"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (c *Client) CreateAccount(ctx context.Context, account *account.Account) (*account.Account, error) {
	ac := &protobuff.Account{
		ID:        account.ID().String(),
		Title:     account.Title().String(),
		Login:     account.Login(),
		Password:  account.Password(),
		URL:       account.URL(),
		Comment:   account.Comment(),
		CreatedAt: account.CreatedAt().Format(time.RFC3339),
		UpdatedAt: account.UpdatedAt().Format(time.RFC3339),
	}

	req := &protobuff.CreateAccountRequest{Data: ac}

	res, err := c.KeeperServiceClient.CreateAccount(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, repository.ErrDataNotFound
	}

	return account, nil
}

func (c *Client) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	req := &protobuff.DeleteAccountRequest{
		ID: id.String(),
	}

	_, err := c.KeeperServiceClient.DeleteAccount(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateAccount(ctx context.Context, account *account.Account) (*account.Account, error) {
	ac := &protobuff.Account{
		ID:        account.ID().String(),
		Title:     account.Title().String(),
		Login:     account.Login(),
		Password:  account.Password(),
		URL:       account.URL(),
		Comment:   account.Comment(),
		CreatedAt: account.CreatedAt().Format(time.RFC3339),
		UpdatedAt: account.UpdatedAt().Format(time.RFC3339),
	}
	req := &protobuff.UpdateAccountRequest{Data: ac}

	res, err := c.KeeperServiceClient.UpdateAccount(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, repository.ErrDataNotFound
	}

	return account, nil
}

func (c *Client) ListAccount(ctx context.Context, parameter queryParameter.QueryParameter) (*account.ListAccountViewModel, error) {

	req := &protobuff.ListAccountRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	res, err := c.KeeperServiceClient.ListAccount(ctx, req)
	if err != nil {
		return nil, err
	}

	listData, err := dto.FromListAccountResponseToDto(res)
	if err != nil {
		return nil, err
	}

	list := &account.ListAccountViewModel{}

	data, err := dto.ToDomainAccounts(*listData)
	if err != nil {
		return nil, err
	}

	list.Data = data
	res.Limit = res.GetLimit()
	res.Offset = res.GetOffset()
	res.Total = res.GetTotal()

	return list, nil

}

func (c *Client) CountAccount(_ context.Context, _ queryParameter.QueryParameter) (uint64, error) {
	return 0, nil
}
