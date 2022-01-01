package service

import (
	"context"
	"errors"
	_ "google.golang.org/grpc/encoding/gzip"
	"grpc-easy/concrete"
	pb "grpc-easy/ecommerce"
	errs "grpc-easy/error"
	"io"
)

func (m *Manage) GetOrder(ctx context.Context, orid *pb.OrderID) (*pb.Order, error) {

	var o *pb.Order
	orderid := orid.GetValue()
	if orderid > m.OderId || (orderid == m.OderId && m.OrderMap[orderid] == nil) {
		return nil, errs.ErrInvalid("GetOrder.ID", concrete.ConcreteOrderId)
	}
	temp := m.OrderMap[orderid]

	if orderid == m.OderId {
		o = &pb.Order{
			Id:    orderid,
			Items: temp,
			Price: temp[len(temp)-1].Price,
		}
	} else {
		o = &pb.Order{
			Id:          orderid,
			Items:       temp,
			Price:       temp[len(temp)-1].Price,
			Description: temp[len(temp)-1].Description,
			Destination: temp[len(temp)-1].Name,
		}
	}
	return o, nil
}

func (m *Manage) SearchOrder(pro *pb.ProductMsg, stream pb.Manage_SearchOrderServer) error {
	var lenth int64 = 0
	if m.OrderMap[m.OderId] == nil {
		lenth = m.OderId
	} else {
		lenth = m.OderId + 1
	}

	for u := int64(0); u < lenth; u++ {
		for _, pros := range m.OrderMap[u] {
			if pros.Id == pro.Id {
				var o *pb.Order
				temp := m.OrderMap[u]
				if u == m.OderId {
					o = &pb.Order{
						Id:    u,
						Items: temp,
					}
				} else {
					o = &pb.Order{
						Id:          u,
						Items:       temp,
						Price:       temp[len(temp)-1].Price,
						Description: temp[len(temp)-1].GetDescription(),
						Destination: temp[len(temp)-1].GetName(),
					}
				}
				err := stream.Send(o)
				if err != nil {
					return errs.ErrInternal("SearchOrder.stream.send", concrete.ConcreteSend)
				}
				break
			}
		}
	}
	return nil
}

func (m *Manage) DeleteOrder(ctx context.Context, orderid *pb.OrderID) (*pb.Respond, error) {
	ord := orderid.GetValue()
	if ord > m.OderId || (ord == m.OderId && m.OrderMap[ord] == nil) {
		return &pb.Respond{Ok: false, Statue: 0}, errs.ErrInvalid("DeleteOrder.ID", concrete.ConcreteOrderId)
	}

	delete(m.OrderMap, orderid.GetValue())

	if ord < m.OderId {
		return &pb.Respond{Ok: true, Statue: 0}, nil
	} else {
		return &pb.Respond{Ok: true, Statue: 1}, nil
	}
}

func (m *Manage) SureSend(ctx context.Context, su *pb.SureMsg) (*pb.Respond, error) {
	var total float32 = 0
	ord := su.Orderid
	if ord > m.OderId || (ord == m.OderId && m.OrderMap[ord] == nil) {
		return &pb.Respond{Ok: false, Statue: 0}, errs.ErrInvalid("SureSend.ID", concrete.ConcreteOrderId)
	} else if ord < m.OderId {
		return &pb.Respond{Ok: false, Statue: 0}, errs.ErrInvalid("SureSend.ID", concrete.ConcreteSendOrderId)
	}

	for _, pro := range m.OrderMap[ord] {
		total += pro.Price
	}

	o := &pb.ProductMsg{
		Id:          -1,
		Price:       total,
		Description: su.Description,
		Name:        su.Destination,
	}
	m.OrderMap[ord] = append(m.OrderMap[ord], o)

	u := &pb.Order{
		Id:          ord,
		Items:       m.OrderMap[m.OderId],
		Price:       total,
		Description: su.Description,
		Destination: su.Destination,
	}
	_, ok := m.Parcel[su.Destination]
	if !ok {
		m.Parcel[su.Destination] = make([]*pb.Order, 0, 10)
	}
	m.Parcel[su.Destination] = append(m.Parcel[su.Destination], u)

	m.OderId++
	m.OrderMap[m.OderId] = make([]*pb.ProductMsg, 0, 10)
	return &pb.Respond{Ok: true, Statue: 1}, nil
}

func (m *Manage) AddOrder(stream pb.Manage_AddOrderServer) error {
	for {
		pro, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				o := &pb.Order{
					Id:    m.OderId,
					Items: m.OrderMap[m.OderId],
				}
				err = stream.SendAndClose(o)
				if err != nil {
					return errs.ErrInternal("AddOrder.stream.SendAndClose", concrete.ConcreteSend)
				}
				break
			} else {
				return errs.ErrInternal("AddOrder.stream.Recv", concrete.ConcreteSend)
			}
		}
		m.OrderMap[m.OderId] = append(m.OrderMap[m.OderId], pro)
	}
	return nil
}
