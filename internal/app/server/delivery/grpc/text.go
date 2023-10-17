package grpc

import (
	"context"
	"time"

	converterText "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/text"
	domainText "github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/filter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateText creating a new text
func (d *Delivery) CreateText(ctx context.Context, req *protobuff.CreateTextRequest) (*protobuff.CreateTextResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "text data authorization error: %v", err)
	}

	titleObj, err := title.New(req.Data.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data title validation error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, req.Data.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, req.Data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	dText, err := domainText.NewWithID(
		id,
		userID,
		*titleObj,
		req.Data.Data,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucText.Create(ctx, dText)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data create error: %v", err)
	}

	return converterText.ToCreateTextResponse(res), nil
}

// UpdateText updating the text
func (d *Delivery) UpdateText(ctx context.Context, req *protobuff.UpdateTextRequest) (*protobuff.UpdateTextResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "text data authorization error: %v", err)
	}

	titleObj, err := title.New(req.Data.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data title validation error: %v", err)
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data comment validation error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, req.Data.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, req.Data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	dText, err := domainText.NewWithID(
		id,
		userID,
		*titleObj,
		req.Data.Data,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucText.Update(ctx, dText)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data update error: %v", err)
	}

	return converterText.ToUpdateTextResponse(res), nil
}

// DeleteText deleting text
func (d *Delivery) DeleteText(ctx context.Context, req *protobuff.DeleteTextRequest) (*protobuff.DeleteTextResponse, error) {
	id := converter.StringToUUID(req.GetID())

	err := d.ucText.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data update error: %v", err)
	}

	return converterText.ToDeleteTextResponse(id), nil
}

// ListText get the test list
func (d *Delivery) ListText(ctx context.Context, req *protobuff.ListTextRequest) (*protobuff.ListTextResponse, error) {
	var parameter queryParameter.QueryParameter
	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "text data authorization error: %v", err)
	}

	parameter.Filters = []filter.Filter{
		{Key: "user_id", Value: userID},
		{Key: "is_deleted", Value: false},
	}
	parameter.Pagination.Limit = req.Limit
	parameter.Pagination.Offset = req.Offset

	res, err := d.ucText.List(ctx, parameter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "text data list error: %v", err)
	}

	return converterText.ToListTextResponse(res), nil
}
