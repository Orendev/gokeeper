package interceptors

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/configs"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthInterceptor is a client interceptor for authentication
type AuthInterceptor struct {
	accessibleRoles map[string][]string
	options         *configs.ServerGRPC
	accessToken     string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(
	accessibleRoles map[string][]string,
	options *configs.ServerGRPC,
) (*AuthInterceptor, error) {

	interceptor := &AuthInterceptor{
		accessibleRoles: accessibleRoles,
		options:         options,
	}

	return interceptor, nil
}

func (interceptor *AuthInterceptor) SetToken(token string) bool {
	if interceptor.accessToken != token {
		interceptor.accessToken = token
		return true
	}
	return false
}

// UnaryAuth returns a client interceptor to authenticate unary RPC
func (interceptor *AuthInterceptor) UnaryAuth() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		logger.Log.Info("--> ",
			zap.Any("unary interceptor: ", method),
		)

		_, ok := interceptor.accessibleRoles[method]
		if ok {
			// everyone can access
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, auth.AuthorizationKey, interceptor.accessToken)
}
