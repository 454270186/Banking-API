package main

import (
	"RESTful/domain"
	"RESTful/global"
	"RESTful/initial"
	"RESTful/logger"
	"fmt"
	"log"
)

var port string

func init() {
	initial.InitialConfig()
	port = fmt.Sprintf(":%d", global.Settings.Port)
}

func main() {
	DB, err := domain.NewDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	r := router(DB)

	logger.Info("Starting listening on port 8080")
	log.Printf("%s starting listening on port %s\n", global.Settings.ProgramName, port)
	r.Run(port)
}
