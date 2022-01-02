package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func MangeInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	log.Printf(info.FullMethod)

	m, err := handler(ctx, req)

	log.Printf("Post %v", m)
	return m, err
}
