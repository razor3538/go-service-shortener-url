package main

import (
	"example.com/m/v2/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env.example file")
	}

	address := os.Getenv("SERVER_ADDRESS")

	if address == "" {
		address = "localhost:8080"
	}

	r := routes.SetupRouter()

	if err := r.Run(address); err != nil {
		panic(err)
	}
}
