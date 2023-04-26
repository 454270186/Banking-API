package main

import (
	"RESTful/domain"
	"RESTful/handlers"
	"RESTful/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)


func router(db *sql.DB) *gin.Engine {
	router := gin.Default()

	ch := handlers.CustomerHandler{
		Service: service.NewCustomerService(domain.NewCustomerRepoDB(db)),
	}
	ah := handlers.AccountHandler{
		Service: service.NewAccountService(domain.NewAccountRepoDB(db)),
	}
	am := AuthMiddleWare{
		repo: domain.NewAuthRepo(),
	}

	router.Use(am.AuthMid())

	router.GET("/", handlers.Hello)

	router.GET("/customer", ch.GetAllCustomer)
	router.GET("/customer/:id", ch.GetCustomerByID)
	router.POST("/customer/:id/account", ah.NewAccount) // create a new account with given id
	router.POST("/customer/:id/account/:account_id", ah.MakeTransaction)

	// time api
	rTime := router.Group("/api")
	{
		rTime.GET("/time", handlers.GetTime)
	}


	return router
}