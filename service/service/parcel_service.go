package service

import (
	"errors"
	_ "google.golang.org/grpc/encoding/gzip"
	pb "grpc-easy/ecommerce"
	errs "grpc-easy/errs"
	"io"
)

func (m *Manage) ShowParcel(stream pb.Manage_ShowParcelServer) error {
	for {
		des, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return errs.ErrInternal("ShowParcel.stream.Recv", errs.ConcreteSend)
			}
		}

		value, ok := m.Parcel[des.GetDes()]
		if !ok {
			return errs.ErrNotFind("ShowParcel.Des", errs.ConcreteDes)
		}
		p := &pb.Parcel{
			Des:    des.GetDes(),
			Count:  int64(len(m.Parcel[des.GetDes()])),
			Orders: value,
		}
		err = stream.Send(p)
		if err != nil {
			return errs.ErrInternal("ShowParcel.stream.send", errs.ConcreteSend)
		}
	}
	return nil
}
