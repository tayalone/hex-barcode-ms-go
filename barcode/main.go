package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/service"
	"github.com/tayalone/hex-barcode-ms-go/barcode/db/store"
	"github.com/tayalone/hex-barcode-ms-go/barcode/db/store/barcoderepo"
	"github.com/tayalone/hex-barcode-ms-go/barcode/mq"
	"github.com/tayalone/hex-barcode-ms-go/barcode/mq/reqbarcode"
	"github.com/tayalone/hex-barcode-ms-go/barcode/mq/resbarcode"
	"github.com/tayalone/hex-barcode-ms-go/barcode/router"
)

func main() {
	/* ------------- pre defined -------------- */
	myMq := mq.ConnectMQ()

	myConn := myMq.GetConn()
	defer myConn.Close()

	myCh := myMq.GetCh()
	defer myCh.Close()

	qReqBarcode, _ := myMq.CreateQueue(mq.QueueName["req_barcode"])
	qResBarcode, _ := myMq.CreateQueue(mq.QueueName["res_barcode"])

	/* ---------------------------------------- */
	/* ------------- defined API -------------- */
	rdbStore := store.New()
	store := rdbStore.GetInstant()
	barcodeRepo := barcoderepo.New(store)
	resBarcodePb := resbarcode.NewPublisher(myCh, *qResBarcode)
	/* ---------------------------------------- */
	/* ----------- defined Services ---------- */
	barcodeSrv := service.New(barcodeRepo, resBarcodePb)
	/* --------------------------------------- */
	/* ----------- defined SPI --------------- */
	reqBCReceiver := reqbarcode.NewReceiver(myCh, *qReqBarcode, barcodeSrv)
	go reqBCReceiver.Receive()

	routeHandler := router.New(barcodeSrv)
	/* --------------------------------------- */

	r := gin.Default()
	r.GET("/healtz", func(c *gin.Context) {
		storeStatus := myMq.GetStatus()
		if storeStatus == true {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
			return
		}
		c.JSON(http.StatusTeapot, gin.H{
			"Message": "Services Not Ready",
		})
		return
	})

	r.GET("/", routeHandler.FindAll)
	r.GET("/:id", routeHandler.GetByID)
	r.POST("/", routeHandler.Create)
	r.PATCH("/:id", routeHandler.UpdateByID)
	r.DELETE("/:id", routeHandler.DeleteByID)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
