package orderrepo

import (
	"github.com/tayalone/hex-barcode-ms-go/order/core/domain"
	"github.com/tayalone/hex-barcode-ms-go/order/core/dto"
	"gorm.io/gorm"
)

/*OrderRepo is Definition of Value */
type OrderRepo struct {
	db *gorm.DB
}

var orderRepo = OrderRepo{}

/*New do Create Rdb Connection*/
func New(db *gorm.DB) *OrderRepo {
	orderRepo.db = db
	return &orderRepo
}

/*Create New Order */
func (o *OrderRepo) Create(c dto.OrderInput) (domain.Ord, error) {
	orderInst := domain.GetTableStruct(c.CourierCode, c.IsCod)
	r := o.db.Create(orderInst)

	if r.Error != nil {
		return domain.Ord{}, r.Error
	}

	return orderInst.GetData(), nil
}
