package main

import (
	"RESTful/domain"
	"RESTful/handlers"
	"RESTful/service"

	"github.com/gin-gonic/gin"
)


func router() *gin.Engine {
	router := gin.Default()

	// ch := handlers.CustomerHandler{
	// 	Service: service.NewCustomerService(domain.NewCustomerRepoStub()),
	// }
	ch := handlers.CustomerHandler{
		Service: service.NewCustomerService(domain.NewCustomerRepoDB()),
	}

	router.GET("/", handlers.Hello)

	router.GET("/customer", ch.GetAllCustomer)
	router.POST("/customer", handlers.CreateCustomer)
	router.GET("/customer/:id", handlers.GetCustomerID)

	// time api
	rTime := router.Group("/api")
	{
		rTime.GET("/time", handlers.GetTime)
	}


	return router
}