package main

import (
	"RESTful/domain"
	"RESTful/logger"
	"os"
	"fmt"
)

var port string

func init() {
	port = os.Getenv("port")
	fmt.Println("port is : " + port)
}

func main() {
	DB, err := domain.NewDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	r := router(DB)

	logger.Info("Starting listening on port 8080")
	r.Run(port)
}
