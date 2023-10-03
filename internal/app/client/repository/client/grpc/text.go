package grpc

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/converter"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// CreateText Let's create a custom post with text
func (c *Client) CreateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {

	ac := &protobuff.Data{
		ID:      text.ID().String(),
		Title:   text.Title().String(),
		Data:    text.Data(),
		Comment: text.Comment(),
	}
	req := &protobuff.CreateTextRequest{Data: ac}

	res, err := c.KeeperServiceClient.CreateText(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, errors.New("error create text data")
	}

	return text, nil
}

// UpdateText Let's update the post with text
func (c *Client) UpdateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {

	ac := &protobuff.Data{
		ID:      text.ID().String(),
		Title:   text.Title().String(),
		Data:    text.Data(),
		Comment: text.Comment(),
	}

	req := &protobuff.UpdateTextRequest{Data: ac}

	res, err := c.KeeperServiceClient.UpdateText(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, errors.New("error update text data")
	}

	return text, nil
}

// DeleteText Delete an arbitrary entry with text
func (c *Client) DeleteText(ctx context.Context, ID uuid.UUID) error {
	req := &protobuff.DeleteTextRequest{
		ID: ID.String(),
	}

	res, err := c.KeeperServiceClient.DeleteText(ctx, req)
	if err != nil {
		return err
	}

	if len(res.GetID()) == 0 {
		return errors.New("error delete text data")
	}

	return nil
}

// ListText Get a list of records with text
func (c *Client) ListText(ctx context.Context, parameter queryParameter.QueryParameter) (*text.ListTextViewModel, error) {
	req := &protobuff.ListTextRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	res, err := c.KeeperServiceClient.ListText(ctx, req)
	if err != nil {
		return nil, err
	}

	listData, err := dto.FromListTextResponseToDto(res)
	if err != nil {
		return nil, err
	}

	list := &text.ListTextViewModel{}

	data, err := converter.ToDomainTexts(*listData)
	if err != nil {
		return nil, err
	}

	list.Data = data
	res.Limit = res.GetLimit()
	res.Offset = res.GetOffset()
	res.Total = res.GetTotal()

	return list, nil
}

// CountText get the number of records
func (c *Client) CountText(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	panic("realize me")
}
