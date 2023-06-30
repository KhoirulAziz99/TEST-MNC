package handler

import (
	"net/http"

	"github.com/KhoirulAziz99/mnc/internal/models"
	"github.com/KhoirulAziz99/mnc/internal/services"
	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	Created(c *gin.Context) 
	FindAll(c *gin.Context) 
	FindByEmail(c *gin.Context)
}

type customerHnadler struct {
	customerServices services.CustomerServices
}

func NewCustomerHandler(customerServices services.CustomerServices) *customerHnadler {
	return &customerHnadler{customerServices: customerServices}
}

func(h customerHnadler) Created(c *gin.Context) {
	var customer models.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message : ":"Failed binding data",
			"error": err,
		})
	}

	err = h.customerServices.Register(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message:":"Failed to register",
			"error":err,
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{"message":"Succsessfully added data"})
}

func (h customerHnadler) FindAll(c *gin.Context) {
	result, err := h.customerServices.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message:":"Failed get all data",
			"error:":err,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data":result})
}

func (h customerHnadler) FindByEmail(c *gin.Context) {
	email := c.Param("email")
	result, err := h.customerServices.GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Failed to Get data",
			"error :": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data":result})
}