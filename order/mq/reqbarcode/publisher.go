package reqbarcode

import (
	"encoding/json"
	"errors"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/tayalone/hex-barcode-ms-go/order/core/dto"
	"github.com/tayalone/hex-barcode-ms-go/order/mq"
)

/*PbInst is a Queue Reciever Instant */
type PbInst struct {
	ch *amqp.Channel
	q  amqp.Queue
}

var myPbInst = PbInst{}

/*NewPublisher do Return Qeue Reciever */
func NewPublisher(ch *amqp.Channel, q amqp.Queue) *PbInst {
	myReInst.ch = ch
	myReInst.q = q
	return &myPbInst
}

/*PushMessage Send Message to Reciever */
func (r PbInst) PushMessage(i dto.PublisherInput) error {
	ch := r.ch

	body, errMs := json.Marshal(i)
	if errMs != nil {
		log.Println("Marshall Error")
		return errors.New("Marshall Error")
	}

	err := ch.Publish(
		"",                          // exchange
		mq.QueueName["req_barcode"], // routing key
		false,                       // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		log.Println("[x] Sent Barcode Fail", err.Error())
		return errors.New(err.Error())
	}
	log.Printf("[x] Sent Payload %+v Success!!\n", i)
	return nil
}
