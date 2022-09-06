package barcoderepo

import (
	"github.com/tayalone/hex-barcode-ms-go/barcode/core"
	"gorm.io/gorm"
)

/*BarcodeRepo is Definition of Value */
type BarcodeRepo struct {
	db *gorm.DB
}

var barcodeRepo = BarcodeRepo{}

/*New do Create Rdb Connection*/
func New(db *gorm.DB) BarcodeRepo {
	barcodeRepo.db = db
	return barcodeRepo
}

/*GetAll Barcode Codition From Repo*/
func (b BarcodeRepo) GetAll() []core.BarcodeCondition {
	var bcs []core.BarcodeCondition
	b.db.Order("created_at desc,id desc").Find(&bcs)
	return bcs
}
