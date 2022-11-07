package main

import (
	"example.com/m/v2/config"
	"flag"

	"example.com/m/v2/routes"
)

func main() {
	config.CheckFlagEnv()
	flag.Parse()

	address := config.Env.Address

	r := routes.SetupRouter()

	if err := r.Run(address); err != nil {
		panic(err)
	}
}
