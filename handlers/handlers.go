package handlers

import (
	"RESTful/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
	Gender  string `json:"gender" xml:"gender"`
}

type CustomerHandler struct {
	Service service.CustomerService
}

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (ch *CustomerHandler) GetAllCustomer(ctx *gin.Context) {
	// customer := []Customer{
	// 	{"xiaofei", "China", "12345", "male"},
	// 	{"erfei", "Chengdu", "6666", "male"},
	// }
	customer, err := ch.Service.GetAllCustomer()
	if err != nil {
		log.Fatal(err)
	}

	if ctx.GetHeader("Content-Type") == "application/json" {
		ctx.JSON(200, customer)
	} else if ctx.GetHeader("Content-Type") == "application/xml" {
		ctx.XML(200, customer)
	}
}

func CreateCustomer(ctx *gin.Context) {
	ctx.String(200, "Post request received...")
}

func GetCustomerID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be a number",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"id": idInt,
	})
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
