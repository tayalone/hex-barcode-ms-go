package reqbarcode

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/ports"
)

/*ReInst is a Queue Reciever Instant */
type ReInst struct {
	ch *amqp.Channel
	q  amqp.Queue
	s  ports.BarcodeSrv
}

var myReInst = ReInst{}

/*NewReceiver do Return Qeue Reciever */
func NewReceiver(ch *amqp.Channel, q amqp.Queue, s ports.BarcodeSrv) ReInst {
	myReInst.ch = ch
	myReInst.q = q
	myReInst.s = s
	return myReInst
}

/*Receive Waiting Message From Queue */
func (r ReInst) Receive() {
	err := r.ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalln("reqbarcode Qos error", err.Error())
	}

	msgs, err := r.ch.Consume(
		QueueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalln("reqbarcode Consume error", err.Error())
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var p dto.ReceiverInput
			err := json.Unmarshal(d.Body, &p)

			if err == nil {
				log.Println("payload is ", p)
			}

			srv := r.s

			errPbBc := srv.PublishBarcode(p)

			if errPbBc != nil {
				log.Println("qResBarcode.Publish error", errPbBc.Error())
			} else {
				log.Println("Done")
				d.Ack(false)
			}

		}
	}()

	log.Println("[*] Waiting for REQ_BARCODE messages")
	<-forever
}
