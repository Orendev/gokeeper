package grpc

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dto"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// CreateBinary Let's create a custom post with binary data
func (c *Client) CreateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {

	ac := &protobuff.Data{
		ID:        binary.ID().String(),
		Title:     binary.Title().String(),
		Data:      binary.Data(),
		Comment:   binary.Comment(),
		CreatedAt: binary.CreatedAt().Format(time.RFC3339),
		UpdatedAt: binary.UpdatedAt().Format(time.RFC3339),
	}
	req := &protobuff.CreateBinaryRequest{Data: ac}

	res, err := c.KeeperServiceClient.CreateBinary(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, repository.ErrDataNotFound
	}

	return binary, nil
}

// UpdateBinary Let's update the post with binary data
func (c *Client) UpdateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {

	ac := &protobuff.Data{
		ID:        binary.ID().String(),
		Title:     binary.Title().String(),
		Data:      binary.Data(),
		Comment:   binary.Comment(),
		CreatedAt: binary.CreatedAt().Format(time.RFC3339),
		UpdatedAt: binary.UpdatedAt().Format(time.RFC3339),
	}

	req := &protobuff.UpdateBinaryRequest{Data: ac}

	res, err := c.KeeperServiceClient.UpdateBinary(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(res.GetID()) == 0 {
		return nil, repository.ErrDataNotFound
	}

	return binary, nil
}

// DeleteBinary Delete an arbitrary entry with binary data
func (c *Client) DeleteBinary(ctx context.Context, ID uuid.UUID) error {
	req := &protobuff.DeleteBinaryRequest{
		ID: ID.String(),
	}

	res, err := c.KeeperServiceClient.DeleteBinary(ctx, req)
	if err != nil {
		return err
	}

	if len(res.GetID()) == 0 {
		return errors.New("error delete binary data")
	}

	return nil
}

// ListBinary Get a list of records with binary data
func (c *Client) ListBinary(ctx context.Context, parameter queryParameter.QueryParameter) (*binary.ListBinaryViewModel, error) {
	req := &protobuff.ListBinaryRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	res, err := c.KeeperServiceClient.ListBinary(ctx, req)
	if err != nil {
		return nil, err
	}

	listData, err := dto.FromListBinaryResponseToDto(res)
	if err != nil {
		return nil, err
	}

	list := &binary.ListBinaryViewModel{}

	data, err := dto.ToDomainBinaries(*listData)
	if err != nil {
		return nil, err
	}

	list.Data = data
	list.Limit = res.GetLimit()
	list.Offset = res.GetOffset()
	list.Total = res.GetTotal()

	return list, nil
}

// CountBinary get the number of records
func (c *Client) CountBinary(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return 0, nil
}
