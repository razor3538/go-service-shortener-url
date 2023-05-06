package main

import (
	"example.com/m/v2/internal/config"
	"example.com/m/v2/internal/routes"
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

	if config.Env.Enable {
		err := r.RunTLS(address, "./testdata/server.pem", "./testdata/server.key")
		if err != nil {
			panic(err)
		}
	} else {
		if err := r.Run(address); err != nil {
			panic(err)
		}
	}
}
