package question

import (
	pb "client/ecommerce"
	errs "client/err"
	"context"
	"log"
)

func addProduct(ctx context.Context, c pb.ManageClient, pro *pb.ProductMsg) {
	r, err := c.AddProduct(ctx, pro)
	if err != nil {
		log.Fatalf("Could not add product")
	}
	log.Printf("AddProduct orderid = %v , productid = %v", r.GetOrderID(), r.GetProductID())
}

func getProduct(ctx context.Context, c pb.ManageClient, in *pb.Information) {
	proMsg, err := c.GetProduct(ctx, in)

	if err != nil {
		errs.ErrDetail(err)
		return
	}
	log.Println("GetProduct R", proMsg)
}

func deleteProduct(ctx context.Context, c pb.ManageClient, in *pb.Information) {
	re, err := c.DeleteProduct(ctx, in)
	if err != nil {
		errs.ErrDetail(err)
		return
	}
	log.Printf("Success %s", re)
}

func getOrder(ctx context.Context, c pb.ManageClient, ordid *pb.OrderID) {
	orders, err := c.GetOrder(ctx, ordid)
	if err != nil {
		errs.ErrDetail(err)
		return
	}

	log.Printf("id :%d , price :%f", orders.Id, orders.Price)
	for _, item := range orders.Items {
		log.Println(item)
	}
}
