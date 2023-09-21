package interceptors

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/useCase/adapters/storage"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthInterceptor is a client interceptor for authentication
type AuthInterceptor struct {
	adapterStorage storage.User
	authMethods    map[string]bool
	accessToken    string
	email          email.Email
	password       password.Password
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(
	adapterStorage storage.User,
	authMethods map[string]bool,
	refreshDuration time.Duration,
	email email.Email,
	password password.Password,
) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		adapterStorage: adapterStorage,
		authMethods:    authMethods,
		email:          email,
		password:       password,
	}

	err := interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil, err
	}

	return interceptor, nil
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

		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, auth.AuthorizationKey, interceptor.accessToken)
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	err := interceptor.refreshToken()
	if err != nil {
		return err
	}

	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()

	return nil
}

func (interceptor *AuthInterceptor) refreshToken() error {

	accessToken, err := interceptor.adapterStorage.LoginUser(context.Background(), interceptor.email, interceptor.password)
	if err != nil {
		return err
	}

	interceptor.accessToken = accessToken.String()

	logger.Log.Info("--> ",
		zap.Any("token refreshed: ", accessToken),
	)

	return nil
}
