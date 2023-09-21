package interceptors

import (
	"context"

	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const servicePath = "/gophkeeper.proto.Keeper/"

type AuthInterceptor struct {
	jwtManager      *auth.JWTManager
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(jwtManager *auth.JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleRoles}
}

func AccessibleRoles() map[string][]string {

	return map[string][]string{
		servicePath + "CreateAccount": {"admin", "user"},
		servicePath + "UpdateAccount": {"user"},
	}
}

func (interceptor *AuthInterceptor) UnaryAuth() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Log.Info("--> unary interceptor: ",
			zap.Any("Full Method", info.FullMethod),
		)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md[auth.AuthorizationKey]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
