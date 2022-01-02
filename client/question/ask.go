package question

import (
	pb "client/ecommerce"
	"context"
)

func Operation(ctx context.Context, c pb.ManageClient) {
	name := "Apple iPhone 11"
	description := "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
	price := float32(699.00)
	pro := &pb.ProductMsg{
		Id:          0,
		Name:        name,
		Description: description,
		Price:       price,
	}
	addProduct(ctx, c, pro)
	pro.Id = 1
	addProduct(ctx, c, pro)

	in := &pb.Information{
		OrderID:   0,
		ProductID: 0,
	}
	getProduct(ctx, c, in)

	deleteProduct(ctx, c, in)

	id := &pb.OrderID{Value: 0}
	getOrder(ctx, c, id)

	searchOrder(ctx, c, pro)
}
