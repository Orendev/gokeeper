package interceptors

import (
	"context"

	"github.com/Orendev/gokeeper/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func (interceptor *AuthInterceptor) UnaryLogger() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		logger.Log.Info("--> unary interceptor: ",
			zap.Any("reg", req),
			zap.Any("Full Method", info.FullMethod),
			zap.Any("Handler", handler),
		)

		return handler(ctx, req)
	}
}
