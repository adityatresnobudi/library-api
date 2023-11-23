package interceptor

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/adityatresnobudi/library-api/internal/shared"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

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
