package main

import (
	"log"

	"github.com/lipesalin/ecom/cmd/api"

)

func main() {
	server := api.APIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
