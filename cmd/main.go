package main

import "RESTful/domain"

func main() {
	DB, err := domain.NewDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	r := router(DB)

	r.Run(":8080")
}
