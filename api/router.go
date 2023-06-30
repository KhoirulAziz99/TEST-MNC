package api

import (
	"database/sql"

	"github.com/KhoirulAziz99/mnc/api/handler"
	"github.com/KhoirulAziz99/mnc/internal/repository"
	"github.com/KhoirulAziz99/mnc/internal/services"
	"github.com/KhoirulAziz99/mnc/pkg/middleware"
	"github.com/gin-gonic/gin"
)


func SetupRouter(db *sql.DB) *gin.Engine {

	customerRepo := repository.NewCustomerRepository(db)
	customerServices := services.NewCustomerServices(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerServices)
	customerLog := handler.NewLogHandler(customerServices)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionServices := services.NewTransactionServices(transactionRepo)
	transactionHadler := handler.NewTransactionHandler(transactionServices)
	
	router := gin.Default() 
	r := router.Group("/customers")
	{
		r.POST("/register", customerHandler.Created)
		r.POST("/login", customerLog.Login)
		r.Use(middleware.AuthMiddleware())
		r.GET("/get-all", customerHandler.FindAll)
		r.POST("/logout", customerLog.Logout)
	}

	r = router.Group("/transaction")
	{	
		r.Use(middleware.AuthMiddleware())
		r.POST("/make", transactionHadler.Create)
		r.GET("/history/:email", transactionHadler.History)
	}

	return router
}