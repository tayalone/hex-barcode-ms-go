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
