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

	router.GET("/", handlers.Hello)

	router.GET("/customer", ch.GetAllCustomer)
	router.GET("/customer/:id", ch.GetCustomerByID)

	// time api
	rTime := router.Group("/api")
	{
		rTime.GET("/time", handlers.GetTime)
	}


	return router
}