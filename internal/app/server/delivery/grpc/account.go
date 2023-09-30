package grpc

import (
	"context"
	"time"

	converterAccount "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/account"
	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/url"
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

	loginObj, err := login.New(req.Data.GetLogin())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account login validation error: %v", err)
	}

	passwordObj, err := password.New(req.Data.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account password validation error: %v", err)
	}

	urlObj, err := url.New(req.Data.GetURL())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account url validation error: %v", err)
	}

	commentObj, err := comment.New(req.Data.GetComment())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account comment validation error: %v", err)
	}

	createdAt := time.Now().UTC()
	updatedAt := createdAt

	dAccount, err := account.NewWithID(
		id,
		userID,
		*titleObj,
		*loginObj,
		*passwordObj,
		*urlObj,
		*commentObj,
		createdAt,
		updatedAt,
	)

	res, err := d.ucAccount.Create(dAccount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "account create error: %v", err)
	}

	return converterAccount.ToCreateAccountResponse(res[0]), nil
}
