package main

import (
	"log"
	"net/http"
	"rpsw/api"
)

func main() {
	api := api.API{}

	// register router
	api.RegisterRoutes()

	port := "8080"
	log.Println("Server listening at: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, api.GetRouter()))
}
