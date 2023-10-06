package grpc

import (
	"context"
	"time"

	converterCard "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/card"
	domainCard "github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/filter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *Delivery) CreateCard(ctx context.Context, req *protobuff.CreateCardRequest) (*protobuff.CreateCardResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "card data authorization error: %v", err)
	}

	createdAt := time.Now().UTC()
	updatedAt := createdAt

	dCard, err := domainCard.NewWithID(
		id,
		userID,
		req.Data.CardNumber,
		req.Data.CardName,
		req.Data.CVC,
		req.Data.CardDate,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucCard.Create(ctx, dCard)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "card data create error: %v", err)
	}

	return converterCard.ToCreateCardResponse(res), nil
}

func (d *Delivery) UpdateCard(ctx context.Context, req *protobuff.UpdateCardRequest) (*protobuff.UpdateCardResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "card data authorization error: %v", err)
	}

	createdAt := time.Now().UTC()
	updatedAt := createdAt

	dCard, err := domainCard.NewWithID(
		id,
		userID,
		req.Data.CardNumber,
		req.Data.CardName,
		req.Data.CVC,
		req.Data.CardDate,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucCard.Update(ctx, dCard)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "card data update error: %v", err)
	}

	return converterCard.ToUpdateCardResponse(res), nil
}

func (d *Delivery) DeleteCard(ctx context.Context, req *protobuff.DeleteCardRequest) (*protobuff.DeleteCardResponse, error) {
	id := converter.StringToUUID(req.GetID())

	err := d.ucCard.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "card data update error: %v", err)
	}

	return converterCard.ToDeleteCardResponse(id), nil
}

func (d *Delivery) ListCard(ctx context.Context, req *protobuff.ListCardRequest) (*protobuff.ListCardResponse, error) {
	var parameter queryParameter.QueryParameter
	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "card data authorization error: %v", err)
	}

	parameter.Filters = []filter.Filter{
		{Key: "user_id", Value: userID},
		{Key: "is_deleted", Value: false},
	}
	parameter.Pagination.Limit = req.Limit
	parameter.Pagination.Offset = req.Offset

	res, err := d.ucCard.List(ctx, parameter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "card data list error: %v", err)
	}

	return converterCard.ToListCardResponse(res), nil
}
