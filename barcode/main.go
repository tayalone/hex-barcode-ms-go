package main

import (
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
	rdbStore := store.New()
	store := rdbStore.GetInstant()

	myMq := mq.ConnectMQ()

	myConn := myMq.GetConn()
	defer myConn.Close()

	myCh := myMq.GetCh()
	defer myCh.Close()

	qReqBarcode, _ := reqbarcode.InitQueue(myCh)
	qResBarcode, _ := resbarcode.InitQueue(myCh)

	/* ---------------------------------------- */
	/* ------------- defined API -------------- */
	barcodeRepo := barcoderepo.New(store)
	resBarcodePb := resbarcode.NewPublisher(myCh, *qResBarcode)
	/* ---------------------------------------- */
	/* ----------- defined Services ---------- */
	barcodeSrv := service.New(barcodeRepo, resBarcodePb)
	/* --------------------------------------- */
	/* ----------- defined SPI --------------- */
	reqBCReceiver := reqbarcode.NewReceiver(myCh, *qReqBarcode, barcodeSrv)
	go reqBCReceiver.Receive()

	router.Init()
	/* --------------------------------------- */
}
