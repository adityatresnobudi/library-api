package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)
	if err != nil {
		out := status.Error(status.Code(err), "error")
		return nil, out
	}

	return res, err
}
