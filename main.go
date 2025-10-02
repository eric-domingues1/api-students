package main

import (
	"log"

	api "github.com/darkpk01/api-students/db/Api"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
