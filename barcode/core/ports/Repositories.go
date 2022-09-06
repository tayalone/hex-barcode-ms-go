package ports

import (
	"github.com/tayalone/hex-barcode-ms-go/barcode/core"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
)

// BarcodeRpstr is Definition of Behavior Repository for Domain
type BarcodeRpstr interface {
	GetAll() []core.BarcodeCondition
	GetByID(uint) (core.BarcodeCondition, error)
	Create(dto.BarCodeInput) (core.BarcodeCondition, error)
	UpdateByID(uint, dto.BarCodeUpdate) error
	DeleteByID(uint) error
	GenBarCode(string, bool, uint) error
}
