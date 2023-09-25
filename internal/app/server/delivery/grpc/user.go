package grpc

import (
	"context"
	jsonUser "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/user"
	domainUser "github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/hashedPassword"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/role"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// RegisterUser creating a new user.
func (d *Delivery) RegisterUser(ctx context.Context, request *protobuff.RegisterUserRequest) (*protobuff.RegisterUserResponse, error) {

	idUser := converter.StringToUUID(request.GetID())

	nameUser, err := name.New(request.GetName())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user name validation error: %v", err)
	}

	roleUser, err := role.New(request.GetRole())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user role validation error: %v", err)
	}

	emailUser, err := email.New(request.GetEmail())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user email validation error: %v", err)
	}

	password, err := password.New(request.GetPassword())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	createdAt := time.Now().UTC()
	updatedAt := createdAt

	if err != nil {
		return nil, err
	}

	dUser, err := domainUser.NewWithID(
		idUser,
		*password,
		*emailUser,
		*nameUser,
		*roleUser,
		createdAt,
		updatedAt,
	)

	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user initialization error: %v", err)
	}

	response, err := d.ucUser.Create(dUser)
	if err != nil {
		logger.Log.Error("error login", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user creation error: %v", err)
	}

	res := jsonUser.ToUserResponse(response[0])
	return res, nil
}

// LoginUser user authorization.
func (d *Delivery) LoginUser(ctx context.Context, request *protobuff.LoginUserRequest) (*protobuff.LoginUserResponse, error) {
	emailUser, err := email.New(request.GetEmail())
	if err != nil {
		logger.Log.Error("error login", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user email validation error: %v", err)
	}

	user, err := d.ucUser.Find(*emailUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	hashedPassword, err := hashedPassword.New(request.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(*hashedPassword) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := d.jwtManager.Generate(user.ID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &protobuff.LoginUserResponse{AccessToken: token}
	return res, nil
}
