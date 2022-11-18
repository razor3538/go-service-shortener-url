package main

import (
	"example.com/m/v2/config"
	"example.com/m/v2/routes"
)

func main() {
	config.CheckFlagEnv()
	config.InitBD()

	address := config.Env.Address

	r := routes.SetupRouter()

	if err := r.Run(address); err != nil {
		panic(err)
	}
}
