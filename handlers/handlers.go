package handlers

import (
	"RESTful/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (ch *CustomerHandler) GetAllCustomer(ctx *gin.Context) {
	customer, appError := ch.Service.GetAllCustomer()
	if appError != nil {
		ctx.JSON(appError.Code, appError.AsMessage())
	}

	if ctx.GetHeader("Content-Type") == "application/json" {
		ctx.JSON(200, customer)
	} else if ctx.GetHeader("Content-Type") == "application/xml" {
		ctx.XML(200, customer)
	}
}

func (ch *CustomerHandler) GetCustomerByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be a number",
		})
		return
	}

	customer, appError := ch.Service.GetCustomerById(id)
	if appError != nil {
		//ctx.String(appError.Code, appError.Message)
		ctx.JSON(appError.Code, appError.AsMessage())
	} else {
		ctx.JSON(200, customer)
	}
}

func GetTime(ctx *gin.Context) {
	timezone := ctx.Query("tz")
	if timezone == "" {
		curTime := time.Now().UTC()
		ctx.JSON(200, gin.H{
			"current_time": curTime,
		})

		return
	}

	response := make(map[string]string)
	tzs := strings.Split(timezone, ",")
	for _, tz := range tzs {
		fmt.Println(tz)
		loc, err := time.LoadLocation(tz)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "invalid timezone",
			})

			return
		}

		now := time.Now().In(loc)

		response[tz] = now.String()
	}

	ctx.JSON(200, response)
}
