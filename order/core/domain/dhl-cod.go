package domain

/*
DHLCod is db schema tmp 'couier_order_dhl'`
use `Composition Technic` https://golangbot.com/inheritance/
*/
type DHLCod struct {
	Ord
}

// GetID get pk
func (c *DHLCod) GetID() uint {
	return c.ID
}

// GetData get Tmp
func (c *DHLCod) GetData() Ord {
	return Ord{
		ID:        c.ID,
		Barcode:   c.Barcode,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

// SetBarcode set initil Barcode
func (c *DHLCod) SetBarcode(barcode string) {
	c.Barcode = &barcode
}
