package oauth

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-easy/errs"
	"log"
)

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	md, ok := metadata.FromIncomingContext(w.ServerStream.Context())
	if !ok {
		return errs.ErrMissMetadata
	}

	if !valid(md["authorization"]) {
		return errs.ErrInvalidToken
	}

	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func EnsureStreamAuth(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println(info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}
