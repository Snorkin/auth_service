package interceptor

import (
	"context"
	"github.com/Snorkin/auth_service/config"
	"github.com/Snorkin/auth_service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type Interceptor struct {
	logger logger.Logger
	cfg    *config.Config
}

func CreateInterceptor(logger logger.Logger, cfg *config.Config) *Interceptor {
	return &Interceptor{
		logger: logger,
		cfg:    cfg,
	}
}

func (i *Interceptor) Log(ctx context.Context, req interface{}, grpcInfo *grpc.UnaryServerInfo, grpcHandler grpc.UnaryHandler) (res interface{}, err error) {
	start := time.Now()
	meta, _ := metadata.FromIncomingContext(ctx)
	res, err = grpcHandler(ctx, req)
	i.logger.Infof("Method: %s, Time: %v, Err: %v", grpcInfo.FullMethod, time.Since(start), meta, err)

	return res, err
}
