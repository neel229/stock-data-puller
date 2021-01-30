package main

import (
	"log"

	"github.com/neel229/stockmarket-simulator/backend/api"
)

func main() {
	server := api.NewServer()
	server.Routes()
	log.Println("Starting the server on port: ")
	server.Start("3000")
}
