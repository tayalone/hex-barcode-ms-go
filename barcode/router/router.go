package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/hex-barcode-ms-go/barcode/db/store"
)

/*Init Http Router */
func Init() {
	r := gin.Default()
	r.GET("/healtz", func(c *gin.Context) {
		storeStatus := store.Status()
		if storeStatus == nil {
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
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
