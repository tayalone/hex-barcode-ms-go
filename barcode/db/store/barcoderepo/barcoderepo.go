package barcoderepo

import (
	"errors"
	"log"

	"github.com/tayalone/hex-barcode-ms-go/barcode/core"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
	"gorm.io/gorm"
)

/*BarcodeRepo is Definition of Value */
type BarcodeRepo struct {
	db *gorm.DB
}

var barcodeRepo = BarcodeRepo{}

var emptyBc = core.BarcodeCondition{}

/*New do Create Rdb Connection*/
func New(db *gorm.DB) BarcodeRepo {
	barcodeRepo.db = db
	return barcodeRepo
}

/*GetAll Barcode Codition From Repo*/
func (b BarcodeRepo) GetAll() []core.BarcodeCondition {
	var bcs []core.BarcodeCondition
	b.db.Order("created_at desc,id desc").Order("created_at desc,id desc").Find(&bcs)
	return bcs
}

/*GetByID Barcode Codition From ID From Repo*/
func (b BarcodeRepo) GetByID(id uint) (core.BarcodeCondition, error) {
	var bc core.BarcodeCondition

	r := b.db.First(&bc, id)

	if r.RowsAffected != 1 {
		return emptyBc, errors.New("Barcode Condition Not Found")
	}
	return bc, nil
}

/*GetByCond return Condition by Courier Code & IsCod From Repo*/
func (b BarcodeRepo) GetByCond(courierCode string, isCod bool) (core.BarcodeCondition, error) {
	var ltBc core.BarcodeCondition
	ltBc.CourierCode = courierCode
	ltBc.IsCod = isCod

	rLtBc := b.db.Order("created_at desc, updated_at desc").Where(&ltBc).First(&ltBc)

	if rLtBc.RowsAffected != 1 {
		return emptyBc, errors.New("Must Be Previos Barcode Condition")
	}
	return ltBc, nil
}

/*Create New Barcode Condition*/
func (b BarcodeRepo) Create(i dto.BarCodeInput) (core.BarcodeCondition, error) {
	ltBc, errLtBc := b.GetByCond(i.CourierCode, i.IsCod)
	if errLtBc != nil {
		return emptyBc, errLtBc
	}
	nBc := core.BarcodeCondition{
		CourierCode:   i.CourierCode,
		IsCod:         i.IsCod,
		StartBarcode:  i.StartBarcode,
		BatchSize:     i.BatchSize,
		PrevCondLogID: ltBc.CondLogID,
		CondLogID:     ltBc.CondLogID + uint(i.BatchSize),
	}

	r := b.db.Create(&nBc)
	if r.Error != nil {
		log.Println("Create Barcode Error, Message", r.Error.Error())
		return emptyBc, errors.New("Create BarCode Error")
	}

	return nBc, nil
}
