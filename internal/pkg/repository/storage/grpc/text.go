package grpc

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dto"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// CreateText Let's create a custom post with text
func (c *Client) CreateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {

	ac := &protobuff.Data{
		ID:        text.ID().String(),
		Title:     text.Title().String(),
		Data:      text.Data(),
		Comment:   text.Comment(),
		CreatedAt: text.CreatedAt().Format(time.RFC3339),
		UpdatedAt: text.UpdatedAt().Format(time.RFC3339),
	}
	req := &protobuff.CreateTextRequest{Data: ac}

	res, err := c.KeeperServiceClient.CreateText(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, repository.ErrDataNotFound
	}

	return text, nil
}

// UpdateText Let's update the post with text
func (c *Client) UpdateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {

	ac := &protobuff.Data{
		ID:        text.ID().String(),
		Title:     text.Title().String(),
		Data:      text.Data(),
		Comment:   text.Comment(),
		CreatedAt: text.CreatedAt().Format(time.RFC3339),
		UpdatedAt: text.UpdatedAt().Format(time.RFC3339),
	}

	req := &protobuff.UpdateTextRequest{Data: ac}

	res, err := c.KeeperServiceClient.UpdateText(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, repository.ErrDataNotFound
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

	data, err := dto.ToDomainTexts(*listData)
	if err != nil {
		return nil, err
	}

	list.Data = data
	list.Limit = res.GetLimit()
	list.Offset = res.GetOffset()
	list.Total = res.GetTotal()

	return list, nil
}

// CountText get the number of records
func (c *Client) CountText(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return 0, nil
}
