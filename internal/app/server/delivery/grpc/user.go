package grpc

import (
	"context"

	jsonUser "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/user"
	domainUser "github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/internal/app/server/domain/user/name"
	"github.com/Orendev/gokeeper/internal/app/server/domain/user/patronymic"
	"github.com/Orendev/gokeeper/internal/app/server/domain/user/surname"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/services/server/internal/domain/user/hashedPassword"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUser creating a new user.
func (d *Delivery) CreateUser(ctx context.Context, request *protobuff.CreateUserRequest) (*protobuff.CreateUserResponse, error) {

	nameUser, err := name.New(request.GetName())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user name validation error: %v", err)
	}

	surnameUser, err := surname.New(request.GetSurname())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user surname validation error: %v", err)
	}

	patronymicUser, err := patronymic.New(request.GetPatronymic())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user patronymic validation error: %v", err)
	}

	emailUser, err := email.New(request.GetEmail())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user email validation error: %v", err)
	}

	hashedPassword, err := hashedPassword.New(request.GetPassword())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	dUser, err := domainUser.New(
		*hashedPassword,
		*emailUser,
		*nameUser,
		*surnameUser,
		*patronymicUser,
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

	if user == nil || !user.IsCorrectPassword(request.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := d.jwtManager.Generate(user.ID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &protobuff.LoginUserResponse{AccessToken: token}
	return res, nil
}
