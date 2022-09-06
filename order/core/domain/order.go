package domain

import (
	"time"
)

// Order create behavior of Courier Order
type Order interface {
	GetID() uint
	GetData() Ord
	SetBarcode(string)
}

// Ord is db schema tmp 'couier_order_CC_ISCOD'`
type Ord struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Barcode   *string   `json:"barcode,omitempty"`
	CreatedAt time.Time `gorm:"index;autoCreateTime"`
	UpdatedAt time.Time `gorm:"index;autoUpdateTime" `
}

// GetTableName from Condition
func GetTableName(courierCode string, isCod bool) string {
	if courierCode == "DHL" {
		if isCod {
			return "order_dhl_cods"
		}
		return "order_oder_dhls"
	}
	return ""
}

// GetTableStruct return emptyu Struct
func GetTableStruct(courierCode string, isCod bool) Order {
	if courierCode == "DHL" {
		if isCod {
			return &DHLCod{}
		}
		return &DHL{}
	}
	return nil
}
