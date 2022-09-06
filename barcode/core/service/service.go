package core

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tayalone/hex-barcode-ms-go/barcode/core"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/ports"
)

/*Service Which Provide Business Logic In Application*/
type Service struct {
	barcodeRpstr ports.BarcodeRpstr
	barcodePb    ports.BarcodePb
}

var srv = Service{}

/*New is Return Ptr of Services*/
func New(r ports.BarcodeRpstr, p ports.BarcodePb) *Service {
	srv.barcodeRpstr = r
	srv.barcodePb = p
	return &srv
}

/*GetAll return Slice of BarcodeCondition from Rpstr*/
func (s *Service) GetAll() []core.BarcodeCondition {
	return s.barcodeRpstr.GetAll()
}

/*GetByID return BarcodeCondition from Rpstr*/
func (s *Service) GetByID(id uint) (core.BarcodeCondition, error) {
	bc, err := s.barcodeRpstr.GetByID(id)
	if err != nil {
		return core.BarcodeCondition{}, err
	}
	return bc, nil
}

/*Create return New Barcode */
func (s *Service) Create(i dto.BarCodeInput) (core.BarcodeCondition, error) {
	bc, err := s.barcodeRpstr.Create(i)
	if err != nil {
		return core.BarcodeCondition{}, err
	}
	return bc, nil
}

/*UpdateByID  return return error whene value error */
func (s *Service) UpdateByID(id uint, u dto.BarCodeUpdate) error {
	err := s.barcodeRpstr.UpdateByID(id, u)
	if err != nil {
		return err
	}
	return nil
}

/*DeleteByID Remove Specific by ID */
func (s *Service) DeleteByID(id uint) error {
	err := s.barcodeRpstr.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

/*GenBarCode Create Barcode By Condition */
func (s *Service) GenBarCode(courierCode string, isCod bool, id uint) (string, error) {
	bc, errC := srv.barcodeRpstr.GetByCond(courierCode, isCod)
	if errC != nil {
		return "", errC
	}
	regFmt := "^([A-Z]{3})+([0-9]{8})+(XTH)"

	match, _ := regexp.MatchString(regFmt, bc.StartBarcode)

	if !match {
		return "", fmt.Errorf("Wrong Barcode Format")
	}

	result := regexp.MustCompile(regFmt).FindAllStringSubmatch(bc.StartBarcode, -1)

	prefix := result[0][1]
	body := result[0][2]
	suffix := result[0][3]

	c, _ := strconv.Atoi(body)
	b := c - int(bc.PrevCondLogID)
	barcode := fmt.Sprintf("%s%s%s", prefix, fmt.Sprintf("%08d", int(id)+b), suffix)

	return barcode, nil
}
