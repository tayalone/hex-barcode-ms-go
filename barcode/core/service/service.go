package service

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
func (s *Service) GenBarCode(i dto.ReceiverInput) (string, error) {
	bc, errC := srv.barcodeRpstr.GetByCond(i.CourierCode, i.IsCod)
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
	barcode := fmt.Sprintf("%s%s%s", prefix, fmt.Sprintf("%08d", int(i.ID)+b), suffix)

	return barcode, nil
}

/*PublishBarcode Create Barcode By Condition */
func (s *Service) PublishBarcode(i dto.ReceiverInput) error {
	barcode, err := s.GenBarCode(i)
	if err != nil {
		return err
	}

	errPush := s.barcodePb.PushMessage(dto.PublisherInput{
		CourierCode: i.CourierCode,
		IsCod:       i.IsCod,
		ID:          i.ID,
		Barcode:     barcode,
	})
	if errPush != nil {
		return errPush
	}

	return nil
}
