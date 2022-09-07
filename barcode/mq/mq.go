package mq

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

/*Conn is Definition of Value */
type Conn interface {
	GetConn() *amqp.Connection
	GetCh() *amqp.Channel
	GetStatus() bool
}

/*MQ is Definition of Value */
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
