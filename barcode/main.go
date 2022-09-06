package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/hex-barcode-ms-go/barcode/db/store"
)

func main() {

	rdbStore := store.New()

	rdbStore.GetInstant()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
