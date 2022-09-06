package ports

import (
	"github.com/tayalone/hex-barcode-ms-go/order/core/domain"
	"github.com/tayalone/hex-barcode-ms-go/order/core/dto"
)

// OrderSrv is Definition of Behavior Services which SPI (Driver) can uses
type OrderSrv interface {
	Create(dto.OrderInput) (domain.Ord, error)
	UpdateBarcode(dto.UpdateBarcode) error
}
