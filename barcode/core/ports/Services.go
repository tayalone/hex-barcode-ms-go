package ports

import (
	"github.com/tayalone/hex-barcode-ms-go/barcode/core"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
)

// BarcodeSrv is Definition of Behavior Services which SPI (Driver) can uses
type BarcodeSrv interface {
	GetAll() []core.BarcodeCondition
	GetByID(id uint) (core.BarcodeCondition, error)
	Create(i dto.BarCodeInput) (core.BarcodeCondition, error)
	UpdateByID(id uint, u dto.BarCodeUpdate) error
	DeleteByID(id uint) error
	GenBarCode(i dto.ReceiverInput) (string, error)
	PublishBarcode(i dto.ReceiverInput) error
}
