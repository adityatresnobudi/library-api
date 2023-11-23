package interceptor

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/adityatresnobudi/library-api/logger"
	"github.com/adityatresnobudi/library-api/shared"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)
	if err != nil {
		out := status.Error(status.Code(err), "error")
		return nil, out
	}

	return res, err
}

func valid(authorization []string) (*shared.JWTClaims, bool) {
	if len(authorization) < 1 {
		return nil, false
	}
	header := strings.TrimPrefix(authorization[0], "Bearer ")
	fmt.Println(header)

	token, err := shared.ValidateJWT(header)
	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(*shared.JWTClaims)
	if !ok || !token.Valid {
		return nil, false
	}

	return claims, true
}

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// authentication (token verification)
	if info.FullMethod == "/auth.Auth/Login" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}

	claims, b := valid(md["authorization"])

	ctx = context.WithValue(ctx, "id", claims.ID)

	if !b {
		return nil, ErrInvalidToken
	}

	m, err := handler(ctx, req)
	if err != nil {
		log.Println("RPC failed with error: %v", err)
	}
	return m, err
}
