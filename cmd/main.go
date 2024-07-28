package main

import (
	"log"

	"github.com/crossstack-q/Go_Prod_API/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
