package main

import (
	"example.com/m/v2/config"
	"example.com/m/v2/routes"
	"github.com/gin-contrib/pprof"
)

// main основная точка входа приложения
func main() {
	config.CheckFlagEnv()
	config.InitBD()

	address := config.Env.Address

	r := routes.SetupRouter()
	pprof.Register(r)

	if err := r.Run(address); err != nil {
		panic(err)
	}
}
