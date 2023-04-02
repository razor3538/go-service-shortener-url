package main

import (
	config2 "example.com/m/v2/internal/config"
	"example.com/m/v2/internal/routes"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/autotls"
	"golang.org/x/crypto/acme/autocert"
	"log"
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

	config2.CheckFlagEnv()
	config2.InitBD()

	address := config2.Env.Address

	r := routes.SetupRouter()
	pprof.Register(r)

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("myPrettyHttpsServerForYandex.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	if config2.Env.EnableHTTPS != "" {
		log.Fatal(autotls.RunWithManager(r, &m))
	} else {
		if err := r.Run(address); err != nil {
			panic(err)
		}
	}
}
