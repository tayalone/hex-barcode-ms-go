package mq

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
)

/*QueueName get Queue Name by Key*/
var QueueName = map[string]string{
	"req_barcode": "REQ_BARCODE",
	"res_barcode": "RES_BARCODE",
}

/*Conn is Definition of Value
 */
type Conn interface {
	GetConn() *amqp.Connection
	GetCh() *amqp.Channel
	GetStatus() bool
	CreateQueue(name string) (*amqp.Queue, error)
}

/*Publisher is Bahavior of Message Q*/
type Publisher interface {
	PushMessage(i dto.ReceiverInput) error
}

/*Receive is Bahavior of Message Q*/
type Receive interface {
	Receive()
}

/*MQ is Definition of Value
 */
type MQ struct {
	mq       *amqp.Connection
	ch       *amqp.Channel
	errMsg   string
	errChMsg string
}

var myMq = MQ{
	mq:       nil,
	ch:       nil,
	errMsg:   "",
	errChMsg: "",
}

/*
GetMQInstance make mq connection
*/
func GetMQInstance() *MQ {
	return &myMq
}

/*
ConnectMQ make mq connection
*/
func ConnectMQ() *MQ {
	rdsn := fmt.Sprintf("amqp://%s:%s@%s:%s",
		os.Getenv("RMQ_USR"),
		os.Getenv("RMQ_PASS"),
		os.Getenv("RMQ_HOST"),
		os.Getenv("RMQ_PORT"))

	conn, err := amqp.Dial(rdsn)
	if err != nil {
		log.Fatalln("Connect MQ Error", err.Error())
		myMq.errMsg = err.Error()
	} else {
		log.Println("Connect MQ Success :)")
		myMq.mq = conn
	}

	ch, errCh := conn.Channel()

	if errCh != nil {
		log.Fatalln("Connect MQ Ch Error", errCh.Error())
		myMq.errMsg = errCh.Error()
	} else {
		log.Println("Connect MQ Ch Success :)")
		myMq.ch = ch
	}

	return &myMq
}

/*
GetConn return rmq connection
*/
func (mq *MQ) GetConn() *amqp.Connection {
	if mq.errMsg == "" {
		return myMq.mq
	}
	return nil
}

/*
GetCh return rmq connection
*/
func (mq *MQ) GetCh() *amqp.Channel {
	if mq.errMsg == "" {
		return mq.ch
	}
	return nil
}

/*
GetStatus return rmq status
*/
func (mq *MQ) GetStatus() bool {
	if mq.errMsg != "" {
		return false
	}
	return true
}

/*CreateQueue with Name*/
func (mq *MQ) CreateQueue(name string) (*amqp.Queue, error) {
	q, err := mq.ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
