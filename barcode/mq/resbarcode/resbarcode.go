package resbarcode

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

/*QueueName of ResBarcode */
var QueueName string = "RES_BARCODE"

/*
InitQueue is init RMQ chan
*/
func InitQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		QueueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		// myInst.errMsg = err.Error()
		// log.Println("Connect Queue", QueueName, "Error", err.Error())
		return nil, err
	}
	return &q, nil
}
