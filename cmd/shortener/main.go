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

	r := routes.SetupRouter()

	if err := r.Run(os.Getenv("SERVER_ADDRESS")); err != nil {
		panic(err)
	}
}
