package reqbarcode

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
)

/*PbInst is a Queue Reciever Instant */
type PbInst struct {
	ch *amqp.Channel
	q  amqp.Queue
}

var myPbInst = PbInst{}

/*NewPublisher do Return Qeue Reciever */
func NewPublisher(ch *amqp.Channel, q amqp.Queue) PbInst {
	myReInst.ch = ch
	myReInst.q = q
	return myPbInst
}

/*PushMessage Send Message to Reciever */
func (r PbInst) PushMessage(i dto.PublisherInput) error {
	log.Println("Do Nothing !!")
	return nil
}
