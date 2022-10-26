package main

import (
	vars "example.com/m/v2/init"
	"example.com/m/v2/routes"
	"flag"
	"os"
)

func main() {
	flag.Parse()

	address := os.Getenv("SERVER_ADDRESS")

	if address == "" {
		if *vars.Flag.Address != "" {
			address = *vars.Flag.Address
		} else {
			address = "localhost:8080"
		}
	}

	r := routes.SetupRouter()

	if err := r.Run(address); err != nil {
		panic(err)
	}
}
