package ports

import "github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"

// BarcodePb is Definition of Behavior Publisher for Domain
type BarcodePb interface {
	PushMessage(dto.PublisherInput) error
}
