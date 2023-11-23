package interceptor

import (
	"context"
	"log"
	"time"

	"github.com/adityatresnobudi/library-api/internal/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	ctx = context.WithValue(ctx, "X-Request-Id", uuid.New())

	res, err := handler(ctx, req)
	if err != nil {
		log.Println("RPC failed with error: %v", err)
	}

	finish := time.Now()
	latency := finish.Sub(start)
	method := info.FullMethod
	statusCode := status.Code(err)

	param := map[string]interface{}{
		"id":          ctx.Value("X-Request-Id"),
		"method":      method,
		"latency":     latency,
		"status_code": statusCode,
	}

	log := logger.NewLogger()
	log.Info(param)

	return res, err
}
