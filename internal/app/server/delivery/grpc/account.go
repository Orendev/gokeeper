package grpc

import (
	"context"
	"time"

	converterAccount "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/account"
	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/filter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *Delivery) CreateAccount(ctx context.Context, req *protobuff.CreateAccountRequest) (*protobuff.CreateAccountResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "account authorization error: %v", err)
	}

	titleObj, err := title.New(req.Data.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account title validation error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, req.Data.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, req.Data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	dAccount, err := account.NewWithID(
		id,
		userID,
		*titleObj,
		req.Data.Login,
		req.Data.Password,
		req.Data.URL,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucAccount.Create(ctx, dAccount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account create error: %v", err)
	}

	return converterAccount.ToCreateAccountResponse(res), nil
}

func (d *Delivery) UpdateAccount(ctx context.Context, req *protobuff.UpdateAccountRequest) (*protobuff.UpdateAccountResponse, error) {
	id := converter.StringToUUID(req.Data.GetID())

	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "account authorization error: %v", err)
	}

	titleObj, err := title.New(req.Data.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account title validation error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, req.Data.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, req.Data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	dAccount, err := account.NewWithID(
		id,
		userID,
		*titleObj,
		req.Data.Login,
		req.Data.Password,
		req.Data.URL,
		req.Data.Comment,
		createdAt,
		updatedAt,
	)

	res, err := d.ucAccount.Update(ctx, dAccount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account update error: %v", err)
	}

	return converterAccount.ToUpdateAccountResponse(res), nil
}

func (d *Delivery) DeleteAccount(ctx context.Context, req *protobuff.DeleteAccountRequest) (*protobuff.DeleteAccountResponse, error) {
	id := converter.StringToUUID(req.GetID())

	err := d.ucAccount.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account update error: %v", err)
	}

	return converterAccount.ToDeleteAccountResponse(id), nil
}

func (d *Delivery) ListAccount(ctx context.Context, req *protobuff.ListAccountRequest) (*protobuff.ListAccountResponse, error) {
	var parameter queryParameter.QueryParameter
	userID, err := d.jwtManager.GetAuthIdentifier(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "account authorization error: %v", err)
	}

	parameter.Filters = []filter.Filter{
		{Key: "user_id", Value: userID},
		{Key: "is_deleted", Value: false},
	}
	parameter.Pagination.Limit = req.Limit
	parameter.Pagination.Offset = req.Offset

	res, err := d.ucAccount.List(ctx, parameter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account list error: %v", err)
	}

	return converterAccount.ToListAccountResponse(res), nil
}
