package dto

/*BarCodeInput is Structor of Bady Params */
type BarCodeInput struct {
	CourierCode  string `json:"courierCode" binding:"required"`
	IsCod        bool   `json:"isCod" binding:"required"`
	StartBarcode string `json:"startBarcode" binding:"required"`
	BatchSize    uint32 `json:"batchSize" binding:"required"`
}

/*BarCodeUpdate is Structor of Bady Params */
type BarCodeUpdate struct {
	CourierCode string `json:"courierCode" binding:"required"`
	IsCod       bool   `json:"isCod" binding:"required"`
	BatchSize   uint32 `json:"batchSize" binding:"required"`
}

/*ReceiverInput is Structor of Bady Params */
type ReceiverInput struct {
	CourierCode string `json:"courierCode" binding:"required"`
	IsCod       bool   `json:"isCod" binding:"required"`
	ID          uint   `json:"id" binding:"required"`
}

/*PublisherInput is Structor of Bady Params */
type PublisherInput struct {
	CourierCode string `json:"courierCode" binding:"required"`
	IsCod       bool   `json:"isCod" binding:"required"`
	ID          uint   `json:"id" binding:"required"`
	Barcode     string `json:"barcode" binding:"required"`
}

/*GetIDUri is Get ID From Parmas */
type GetIDUri struct {
	ID uint `uri:"id" binding:"required"`
}
