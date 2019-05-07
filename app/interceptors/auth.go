package interceptors

import (
	"base/config"
	"base/pkg/arrays"
	"base/pkg/crypto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

const TokenHeaderIndex int = 0
const TokenHeaderName = "auth"

func isProtectedMethod(method string) bool {
	return !arrays.InStrArray(config.UnprotectedMethods, method)
}

// Authorization unary interceptor function to handle authorize per RPC call
func AuthInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Skip authorize when GetJWT is requested
	if config.SecurityCfg.DisableTokenAuth || isProtectedMethod(info.FullMethod) {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}

	// Calls the handler
	h, err := handler(ctx, req)
	time.Since(start)
	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Token must be passed in the header/metadata")
	}
	authHeader, ok := md[TokenHeaderName]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}
	token := authHeader[TokenHeaderIndex]
	// validateToken function validates the token
	cerr := crypto.ValidateToken([]byte(token), config.SecurityCfg.TokenSecret)
	if cerr.IsError() {
		return status.Errorf(codes.Unauthenticated, cerr.Error())
	}
	return nil
}
