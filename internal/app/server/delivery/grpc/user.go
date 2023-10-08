package grpc

import (
	"context"
	"time"

	converterUser "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/user"
	domainUser "github.com/Orendev/gokeeper/internal/pkg/domain/user"
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
)

// RegisterUser creating a new user.
func (d *Delivery) RegisterUser(ctx context.Context, req *protobuff.RegisterUserRequest) (*protobuff.RegisterUserResponse, error) {

	idUser := converter.StringToUUID(req.GetID())

	nameUser, err := name.New(req.GetName())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user name validation error: %v", err)
	}

	roleUser, err := role.New(req.GetRole())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user role validation error: %v", err)
	}

	emailUser, err := email.New(req.GetEmail())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user email validation error: %v", err)
	}

	hashPasswordUser, err := hashedPassword.New(req.GetPassword())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	passwordUser, err := password.New(hashPasswordUser.String())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, req.UpdatedAt)
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

	tokenObject := token.New(tok)

	dUser.SetToken(*tokenObject)

	response, err := d.ucUser.Create(context.Background(), dUser)
	if err != nil {
		logger.Log.Error("error login", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user creation error: %v", err)
	}

	return converterUser.ToRegisterUserResponse(response), nil
}

// LoginUser user authorization.
func (d *Delivery) LoginUser(ctx context.Context, req *protobuff.LoginUserRequest) (*protobuff.LoginUserResponse, error) {
	emailUserObject, err := email.New(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user email validation error: %v", err)
	}

	passwordObject, err := password.New(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	user, err := d.ucUser.Login(context.Background(), *emailUserObject, *passwordObject)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect email/password")
	}

	tok, err := d.jwtManager.Generate(user.ID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	tokenObject := token.New(tok)

	user.SetToken(*tokenObject)

	if !d.ucUser.SetToken(ctx, user) {
		return nil, status.Errorf(codes.Internal, "cannot update token user")
	}

	return converterUser.ToLoginUserResponse(user), nil
}
