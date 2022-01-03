package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMissMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken = status.Errorf(codes.Unauthenticated, "invalid credentials")
)
