package main

import (
	"RESTful/domain"
	"RESTful/logger"
)

func main() {
	DB, err := domain.NewDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	r := router(DB)

	logger.Info("Starting listening on port 8080")
	r.Run(":8080")
}
