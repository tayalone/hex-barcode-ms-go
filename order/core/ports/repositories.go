package ports

import (
	"github.com/tayalone/hex-barcode-ms-go/order/core/domain"
	"github.com/tayalone/hex-barcode-ms-go/order/core/dto"
)

// OrderRpstr is Definition of Behavior Repository for Domain
type OrderRpstr interface {
	Create(dto.OrderInput) (domain.Ord, error)
	Update(dto.UpdateBarcode) error
}
