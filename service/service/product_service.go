package service

import (
	"context"
	_ "google.golang.org/grpc/encoding/gzip"
	"grpc-easy/concrete"
	pb "grpc-easy/ecommerce"
	errs "grpc-easy/error"
)

func (m *Manage) init() {
	m.OrderMap = make(map[int64][]*pb.ProductMsg, 10)
	m.Parcel = make(map[string][]*pb.Order, 10)
	m.OderId = 0
	m.OrderMap[m.OderId] = make([]*pb.ProductMsg, 0, 10)
}

func (m *Manage) AddProduct(ctx context.Context, pro *pb.ProductMsg) (*pb.Information, error) {
	if m.OrderMap == nil {
		m.init()
	}
	m.OrderMap[m.OderId] = append(m.OrderMap[m.OderId], pro)
	return &pb.Information{OrderID: m.OderId, ProductID: pro.Id}, nil
}

func (m *Manage) GetProduct(ctx context.Context, info *pb.Information) (*pb.ProductMsg, error) {
	orderid := info.GetOrderID()
	if orderid > m.OderId || (orderid == m.OderId && m.OrderMap[orderid] == nil) {
		return nil, errs.ErrInvalid("DeleteProduct.ID", concrete.ConcreteOrderId)
	}
	proid := info.GetProductID()
	for _, pro := range m.OrderMap[orderid] {
		if pro.Id == proid {
			return pro, nil
		}
	}
	return nil, errs.ErrNotFind("GetProduct.ID", concrete.ConcreteProductId)
}

func (m *Manage) DeleteProduct(ctx context.Context, info *pb.Information) (*pb.Respond, error) {
	var i, lenth int

	orderid := info.GetOrderID()
	if orderid > m.OderId || (orderid == m.OderId && m.OrderMap[orderid] == nil) {
		return &pb.Respond{Ok: false, Statue: 0}, errs.ErrInvalid("DeleteProduct.ID", concrete.ConcreteOrderId)
	}
	proid := info.GetProductID()
	temp := m.OrderMap[orderid]
	lenth = len(temp)

	for i = 0; i < lenth; i++ {
		if temp[i].Id == proid {
			temp = append(temp[:i], temp[i+1:]...)
			m.OrderMap[orderid] = temp
			break
		}
	}
	if i == lenth {
		return &pb.Respond{Ok: false, Statue: 0}, errs.ErrNotFind("DeleteProduct.ID", concrete.ConcreteProductId)
	}

	// Status 1表示删除的是当前没有发出去的订单的商品  0表示删除的是已经发出订单的商品
	if orderid < m.OderId {
		return &pb.Respond{Ok: true, Statue: 0}, nil
	} else {
		return &pb.Respond{Ok: true, Statue: 1}, nil
	}
}
