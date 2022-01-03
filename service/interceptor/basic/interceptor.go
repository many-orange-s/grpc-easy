package basic

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-easy/errs"
	"log"
	"strings"
)

func EnsureValidBasic(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errs.ErrMissMetadata
	}

	if !valid(md["Product"]) {
		return nil, errs.ErrInvalidToken
	}

	log.Printf(info.FullMethod)

	m, err := handler(ctx, req)
	return m, err
}

func valid(auth []string) bool {
	if len(auth) < 1 {
		return false
	}

	token := strings.TrimPrefix(auth[0], "Basic ")
	return token == base64.StdEncoding.EncodeToString([]byte("太阳高高我要起早"))
}
