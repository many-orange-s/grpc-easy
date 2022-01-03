package question

import (
	pb "client/ecommerce"
	errs "client/err"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"io"
	"log"
)

func addProduct(ctx context.Context, c pb.ManageClient, pro *pb.ProductMsg) {
	r, err := c.AddProduct(ctx, pro)
	if err != nil {
		log.Println(err)
		log.Fatalf("Could not add product")
	}
	log.Printf("AddProduct orderid = %v , productid = %v", r.GetOrderID(), r.GetProductID())
}

func getProduct(ctx context.Context, c pb.ManageClient, in *pb.Information) {
	proMsg, err := c.GetProduct(ctx, in, grpc.UseCompressor(gzip.Name))

	if err != nil {
		errs.ErrDetail(err)
		return
	}
	log.Println("GetProduct R", proMsg)
}

func deleteProduct(ctx context.Context, c pb.ManageClient, in *pb.Information) {
	re, err := c.DeleteProduct(ctx, in, grpc.UseCompressor(gzip.Name))
	if err != nil {
		errs.ErrDetail(err)
		return
	}
	log.Printf("Success %s", re)
}

func getOrder(ctx context.Context, c pb.ManageClient, ordid *pb.OrderID) {
	orders, err := c.GetOrder(ctx, ordid, grpc.UseCompressor(gzip.Name))
	if err != nil {
		errs.ErrDetail(err)
		return
	}

	log.Printf("getorder id :%d ", orders.Id)
	for _, item := range orders.Items {
		log.Println("every product is  ", item)
	}
}

func searchOrder(ctx context.Context, c pb.ManageClient, pro *pb.ProductMsg) {
	stream, err := c.SearchOrder(ctx, pro)
	if err != nil {
		errs.ErrDetail(err)
		return
	}

	for {
		ord, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			errs.ErrDetail(err)
			return
		}

		log.Println("searchorder post ", ord)
	}
}

func addOrder(ctx context.Context, c pb.ManageClient, pros []*pb.ProductMsg) {
	stream, err := c.AddOrder(ctx)
	if err != nil {
		errs.ErrDetail(err)
		return
	}

	for _, pro := range pros {
		_ = stream.Send(pro)
	}

	ord, err := stream.CloseAndRecv()
	log.Println(ord)
}
