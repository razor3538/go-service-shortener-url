package main

import (
	"example.com/m/v2/config"
	"example.com/m/v2/routes"
	"fmt"
	"github.com/gin-contrib/pprof"
)

var (
	buildDate    string
	buildVersion string
	buildCommit  string
)

// main основная точка входа приложения
func main() {
	println(fmt.Sprintf("Build version: %s", buildVersion))
	println(fmt.Sprintf("Build date: %s", buildDate))
	println(fmt.Sprintf("Build commit: %s", buildCommit))

	config.CheckFlagEnv()
	config.InitBD()

	address := config.Env.Address

	r := routes.SetupRouter()
	pprof.Register(r)

	if err := r.Run(address); err != nil {
		panic(err)
	}
}
