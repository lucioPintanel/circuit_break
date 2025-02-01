package config

import (
	"log"
	"os"
)

var (
	ApiUrl string
)

func LoadConfig() {
	ApiUrl = os.Getenv("API_URL")
	if ApiUrl == "" {
		log.Fatal("API_URL environment variable not set")
	}
}
