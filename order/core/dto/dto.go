package dto

/*OrderInput is Structor of Bady Params */
type OrderInput struct {
	CourierCode string `json:"courierCode" binding:"required"`
	IsCod       bool   `json:"isCod" binding:"required"`
}

/*UpdateBarcode is Structor of Bady Params */
type UpdateBarcode struct {
	CourierCode string `json:"courierCode" binding:"required"`
	IsCod       bool   `json:"isCod" binding:"required"`
	ID          uint   `json:"id" binding:"required"`
	Barcode     string `json:"barcode" binding:"required"`
}
