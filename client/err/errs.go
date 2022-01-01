package errs

import (
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func ErrDetail(err error) {
	errorCode := status.Code(err)
	if errorCode == codes.NotFound {
		detail(err, errorCode)
	} else if errorCode == codes.Internal {
		detail(err, errorCode)
	} else if errorCode == codes.InvalidArgument {
		detail(err, errorCode)
	} else {
		log.Printf("Unhandled error : %s", errorCode)
	}
}

func detail(err error, errorCode codes.Code) {
	log.Printf("Error : %s", errorCode)
	errorStatus := status.Convert(err)
	for _, d := range errorStatus.Details() {
		switch info := d.(type) {
		case *epb.BadRequest_FieldViolation:
			log.Printf("Error  %s", info)
		default:
			log.Printf("Unexpected error type: %s", info)
		}
	}
}
