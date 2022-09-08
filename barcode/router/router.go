package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/dto"
	"github.com/tayalone/hex-barcode-ms-go/barcode/core/ports"
)

/*Router is Sturcture */
type Router struct {
	barcodeSrv ports.BarcodeSrv
}

/*Handler is Define Reouter Befavior */
type Handler interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

/*New Return New Routner Handler */
func New(b ports.BarcodeSrv) *Router {
	return &Router{
		barcodeSrv: b,
	}
}

/*FindAll is Handler Http Router */
func (r *Router) FindAll(c *gin.Context) {
	bcs := r.barcodeSrv.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"message":           "OK",
		"barCodeConditions": bcs,
	})
}

/*
GetByID is Get Barcode By Id Condition
*/
func (r *Router) GetByID(c *gin.Context) {
	var gi dto.GetIDUri
	if err := c.ShouldBindUri(&gi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	bc, errGetID := r.barcodeSrv.GetByID(gi.ID)

	if errGetID != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errGetID.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "OK",
		"barCodeCondition": bc,
	})
	return
}

/*Create New Barcode */
func (r *Router) Create(c *gin.Context) {
	var input dto.BarCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bc, errCreate := r.barcodeSrv.Create(input)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "OK",
		"barCodeCondition": bc,
	})
}

/*UpdateByID is Update New Data By Id */
func (r *Router) UpdateByID(c *gin.Context) {
	var gi dto.GetIDUri
	if err := c.ShouldBindUri(&gi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var update dto.BarCodeUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.barcodeSrv.UpdateByID(gi.ID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

/*DeleteByID is Remove Data By Id*/
func (r *Router) DeleteByID(c *gin.Context) {
	var gi dto.GetIDUri
	if err := c.ShouldBindUri(&gi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err := r.barcodeSrv.DeleteByID(gi.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
