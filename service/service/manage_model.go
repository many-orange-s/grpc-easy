package service

import pb "grpc-easy/ecommerce"

type Manage struct {
	OrderMap map[int64][]*pb.ProductMsg
	Parcel   map[string][]*pb.Order
	OderId   int64
}
