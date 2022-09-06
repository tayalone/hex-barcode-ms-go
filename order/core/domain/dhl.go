package domain

/*
DHL is db schema tmp 'couier_order_dhl'`
use `Composition Technic` https://golangbot.com/inheritance/
*/
type DHL struct {
	Ord
}

// GetID get pk
func (c *DHL) GetID() uint {
	return c.ID
}

// GetData get Tmp
func (c *DHL) GetData() Ord {
	return Ord{
		ID:        c.ID,
		Barcode:   c.Barcode,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

// SetBarcode set initil Barcode
func (c *Ord) SetBarcode(barcode string) {
	c.Barcode = &barcode
}
