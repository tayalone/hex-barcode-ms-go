package service

import (
	"github.com/tayalone/hex-barcode-ms-go/order/core/domain"
	"github.com/tayalone/hex-barcode-ms-go/order/core/dto"
	"github.com/tayalone/hex-barcode-ms-go/order/core/ports"
)

/*Service Which Provide Business Logic In Application*/
type Service struct {
	orderRpstr   ports.OrderRpstr
	reqBarcodePb ports.ReqBarcodePb
}

var srv = Service{}

/*New is Return Ptr of Services*/
func New(r ports.OrderRpstr, p ports.ReqBarcodePb) *Service {
	srv.orderRpstr = r
	srv.reqBarcodePb = p
	return &srv
}

/*Create New Order*/
func (s *Service) Create(i dto.OrderInput) (domain.Ord, error) {
	order, err := s.orderRpstr.Create(i)
	if err != nil {
		return domain.Ord{}, err
	}
	errReq := s.reqBarcodePb.Publish(i.CourierCode, i.IsCod, order.ID)
	if errReq != nil {
		return domain.Ord{}, err
	}
	return order, nil
}

/*UpdateBarcode Barcode From Reciever*/
func (s *Service) UpdateBarcode(u dto.UpdateBarcode) error {
	err := s.orderRpstr.Update(u)
	if err != nil {
		return err
	}
	return nil
}
