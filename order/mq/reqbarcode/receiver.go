package reqbarcode

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

/*ReInst is a Queue Reciever Instant */
type ReInst struct {
	ch *amqp.Channel
	q  amqp.Queue
}

var myReInst = ReInst{}

/*NewReceiver do Return Qeue Reciever */
func NewReceiver(ch *amqp.Channel, q amqp.Queue) ReInst {
	myReInst.ch = ch
	myReInst.q = q
	return myReInst
}

/*Receive Waiting Message From Queue */
func (r ReInst) Receive() {
	log.Println("Do notthin.")
}
