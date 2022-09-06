package ports

import (
	"github.com/tayalone/hex-barcode-ms-go/barcode/core"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
)

// BarcodeSrv is Definition of Behavior Services which SPI (Driver) can uses
type BarcodeSrv interface {
	GetAll() []core.BarcodeCondition
	GetByID(uint) (core.BarcodeCondition, error)
	Create(dto.BarCodeInput) (core.BarcodeCondition, error)
	UpdateByID(uint, dto.BarCodeUpdate) error
	DeleteByID(uint) error
	GenBarCode(string, bool, uint) (string error)
}
