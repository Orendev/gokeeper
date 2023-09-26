package grpc

import (
	"context"
	converterUser "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/user"
	domainUser "github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/hashedPassword"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/role"
	"github.com/Orendev/gokeeper/pkg/type/token"
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

	hashPasswordUser, err := hashedPassword.New(request.GetPassword())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	passwordUser, err := password.New(hashPasswordUser.String())
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
		*passwordUser,
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

	tok, err := d.jwtManager.Generate(dUser.ID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	tokenObject, err := token.New(tok)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	dUser.SetToken(*tokenObject)

	response, err := d.ucUser.Create(dUser)
	if err != nil {
		logger.Log.Error("error login", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user creation error: %v", err)
	}

	return converterUser.ToRegisterUserResponse(response[0]), nil
}

// LoginUser user authorization.
func (d *Delivery) LoginUser(ctx context.Context, request *protobuff.LoginUserRequest) (*protobuff.LoginUserResponse, error) {
	emailUserObject, err := email.New(request.GetEmail())
	if err != nil {
		logger.Log.Error("error login", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user email validation error: %v", err)
	}

	user, err := d.ucUser.Find(*emailUserObject)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	hashedPasswordObject, err := hashedPassword.New(request.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(*hashedPasswordObject) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	tok, err := d.jwtManager.Generate(user.ID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	tokenObject, err := token.New(tok)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	user.SetToken(*tokenObject)

	return converterUser.ToLoginUserResponse(user), nil
}
