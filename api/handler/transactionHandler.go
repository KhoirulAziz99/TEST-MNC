package handler

import (
	"net/http"

	"github.com/KhoirulAziz99/mnc/internal/models"
	"github.com/KhoirulAziz99/mnc/internal/services"
	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	Create(c *gin.Context)
	History(c *gin.Context)
}
type transactionHandler struct {
	transactionServices services.TransactionServices
}

func NewTransactionHandler(transactionServices services.TransactionServices) *transactionHandler {
	return &transactionHandler{transactionServices: transactionServices}
}

//method untuk membuat transaksi baru
func(t transactionHandler) Create(c *gin.Context) {
	var transaction models.Transaction
	err :=c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Failed binding data transaction",
			"err":err,
		})
	}
	err = t.transactionServices.Make(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Failed make transaction",
			"error":err,
		})
	}
	c.JSON(http.StatusOK, gin.H{"Message":"Successfully added data transaction"})
}

//method untuk menampilkan history transaksi
func(t transactionHandler) History(c *gin.Context) {

	email := c.Param("email")

	history, err :=t.transactionServices.History(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Failed to get history",
			"error":err,
		})
	}
	c.JSON(http.StatusOK, gin.H{"history":history})
}