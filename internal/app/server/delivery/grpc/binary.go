package grpc

import (
	"context"
	"time"

	converterBinary "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/binary"
	domainBinary "github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/filter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBinary creating a new binary
func (d *Delivery) CreateBinary(ctx context.Context, req *protobuff.CreateBinaryRequest) (*protobuff.CreateBinaryResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "binary data authorization error: %v", err)
	}

	titleObj, err := title.New(req.Data.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "binary data title validation error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, req.Data.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, req.Data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	dBinary, err := domainBinary.NewWithID(
		id,
		userID,
		*titleObj,
		req.Data.Data,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucBinary.Create(ctx, dBinary)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data create error: %v", err)
	}

	return converterBinary.ToCreateBinaryResponse(res), nil
}

// UpdateBinary updating the binary
func (d *Delivery) UpdateBinary(ctx context.Context, req *protobuff.UpdateBinaryRequest) (*protobuff.UpdateBinaryResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "binary data authorization error: %v", err)
	}

	titleObj, err := title.New(req.Data.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "binary data title validation error: %v", err)
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "binary data comment validation error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, req.Data.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, req.Data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	dBinary, err := domainBinary.NewWithID(
		id,
		userID,
		*titleObj,
		req.Data.Data,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucBinary.Update(ctx, dBinary)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "binary data update error: %v", err)
	}

	return converterBinary.ToUpdateBinaryResponse(res), nil
}

// DeleteBinary deleting binary
func (d *Delivery) DeleteBinary(ctx context.Context, req *protobuff.DeleteBinaryRequest) (*protobuff.DeleteBinaryResponse, error) {
	id := converter.StringToUUID(req.GetID())

	err := d.ucBinary.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "binary data update error: %v", err)
	}

	return converterBinary.ToDeleteBinaryResponse(id), nil
}

// ListBinary get the binary list
func (d *Delivery) ListBinary(ctx context.Context, req *protobuff.ListBinaryRequest) (*protobuff.ListBinaryResponse, error) {
	var parameter queryParameter.QueryParameter
	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "binary data authorization error: %v", err)
	}

	parameter.Filters = []filter.Filter{
		{Key: "user_id", Value: userID},
		{Key: "is_deleted", Value: false},
	}
	parameter.Pagination.Limit = req.Limit
	parameter.Pagination.Offset = req.Offset

	res, err := d.ucBinary.List(ctx, parameter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "binary data list error: %v", err)
	}

	return converterBinary.ToListBinaryResponse(res), nil
}
