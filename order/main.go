package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/hex-barcode-ms-go/order/mq"
)

func main() {
	/* ------------- pre defined -------------- */
	myMq := mq.ConnectMQ()

	myConn := myMq.GetConn()
	defer myConn.Close()

	myCh := myMq.GetCh()
	defer myCh.Close()

	// qReqBarcode, _ := myMq.CreateQueue(mq.QueueName["req_barcode"])
	// qResBarcode, _ := myMq.CreateQueue(mq.QueueName["res_barcode"])
	// /* ---------------------------------------- */
	// /* ------------- defined API -------------- */
	// rdbStore := store.New()
	// store := rdbStore.GetInstant()
	// orderRepo := orderrepo.New(store)
	// reqBarcodePb := reqbarcode.NewPublisher(myCh, *qReqBarcode)
	// /* ---------------------------------------- */
	// /* -------------- defined Services -------- */
	// orderSrv := service.New(orderRepo, reqBarcodePb)
	// /* --------------------------------------- */

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
