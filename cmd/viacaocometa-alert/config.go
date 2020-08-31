package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	baseURL = mustGetenv("BASE_URL")
	originID = mustGetenv("ORIGIN_ID")
	destinationID = mustGetenv("DESTINATION_ID")
	whatsappRecipient = mustGetenv("WHATSAPP_RECIPIENT")
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}
