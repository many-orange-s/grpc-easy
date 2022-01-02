package main

import (
	pb "client/ecommerce"
	"client/question"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address  = "localhost:50051"
	hostname = "localhost"
	crtFile  = "D:\\go_project\\grpc-easy\\service\\server.crt"
)

func main() {
	/*creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials :%v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	*/
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewManageClient(conn)

	clientDeadline := time.Now().Add(time.Duration(2 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	question.Operation(ctx, c)
}
