package errs

import (
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrNotFind(errs, msg string) error {
	errorStatus := status.New(codes.NotFound, "The requesting entity was not found")
	ds, err := errorStatus.WithDetails(
		&epb.BadRequest_FieldViolation{
			Field:       errs,
			Description: msg,
		},
	)
	if err != nil {
		return errorStatus.Err()
	}
	return ds.Err()
}

func ErrInvalid(errs, msg string) error {
	errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
	ds, err := errorStatus.WithDetails(
		&epb.BadRequest_FieldViolation{
			Field:       errs,
			Description: msg,
		},
	)
	if err != nil {
		return errorStatus.Err()
	}
	return ds.Err()
}

func ErrInternal(errs, msg string) error {
	errorStatus := status.New(codes.Internal, "Internal error")
	ds, err := errorStatus.WithDetails(
		&epb.BadRequest_FieldViolation{
			Field:       errs,
			Description: msg,
		},
	)
	if err != nil {
		return errorStatus.Err()
	}
	return ds.Err()
}
