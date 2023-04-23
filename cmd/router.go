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

	// ch := handlers.CustomerHandler{
	// 	Service: service.NewCustomerService(domain.NewCustomerRepoStub()),
	// }
	ch := handlers.CustomerHandler{
		Service: service.NewCustomerService(domain.NewCustomerRepoDB(db)),
	}
	ah := handlers.AccountHandler{
		Service: service.NewAccountService(domain.NewAccountRepoDB(db)),
	}

	router.GET("/", handlers.Hello)

	router.GET("/customer", ch.GetAllCustomer)
	router.GET("/customer/:id", ch.GetCustomerByID)
	router.POST("/customer/:id([0-9]+)/account", ah.NewAccount) // create a new account with given id

	// time api
	rTime := router.Group("/api")
	{
		rTime.GET("/time", handlers.GetTime)
	}


	return router
}