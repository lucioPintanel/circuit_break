package main

import (
	"circuit-break/internal/client"
	"log"
)

func main() {
	client := client.NewApiClient()
	response, err := client.MakeRequest("https://mockserver/api/endpoint")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	log.Printf("Response: %s", response)
}
