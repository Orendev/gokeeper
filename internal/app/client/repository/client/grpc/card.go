package grpc

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/converter"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// CreateCard Let's create a custom post with text
func (c *Client) CreateCard(ctx context.Context, card *card.CardData) (*card.CardData, error) {

	ac := &protobuff.Card{
		ID:         card.ID().String(),
		CardNumber: card.CardNumber(),
		CardName:   card.CardName(),
		CVC:        card.CVC(),
		CardDate:   card.CardDate(),
		Comment:    card.Comment(),
	}
	req := &protobuff.CreateCardRequest{Data: ac}

	res, err := c.KeeperServiceClient.CreateCard(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, errors.New("error create card data")
	}

	return card, nil
}

// UpdateCard Let's update the post with text
func (c *Client) UpdateCard(ctx context.Context, card *card.CardData) (*card.CardData, error) {

	ac := &protobuff.Card{
		ID:         card.ID().String(),
		CardNumber: card.CardNumber(),
		CardName:   card.CardName(),
		CVC:        card.CVC(),
		CardDate:   card.CardDate(),
		Comment:    card.Comment(),
	}

	req := &protobuff.UpdateCardRequest{Data: ac}

	res, err := c.KeeperServiceClient.UpdateCard(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, errors.New("error update card data")
	}

	return card, nil
}

// DeleteCard Delete an arbitrary entry with text
func (c *Client) DeleteCard(ctx context.Context, ID uuid.UUID) error {
	req := &protobuff.DeleteCardRequest{
		ID: ID.String(),
	}

	res, err := c.KeeperServiceClient.DeleteCard(ctx, req)
	if err != nil {
		return err
	}

	if len(res.GetID()) == 0 {
		return errors.New("error delete card data")
	}

	return nil
}

// ListCard Get a list of records with text
func (c *Client) ListCard(ctx context.Context, parameter queryParameter.QueryParameter) (*card.ListCardViewModel, error) {
	req := &protobuff.ListCardRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	res, err := c.KeeperServiceClient.ListCard(ctx, req)
	if err != nil {
		return nil, err
	}

	listData, err := dto.FromListCardResponseToDto(res)
	if err != nil {
		return nil, err
	}

	list := &card.ListCardViewModel{}

	data, err := converter.ToDomainCards(*listData)
	if err != nil {
		return nil, err
	}

	list.Data = data
	res.Limit = res.GetLimit()
	res.Offset = res.GetOffset()
	res.Total = res.GetTotal()

	return list, nil
}

// CountCard get the number of records
func (c *Client) CountCard(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	panic("realize me")
}
