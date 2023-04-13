package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/autotls"
	"go-service-shortener-url/internal/config"
	"go-service-shortener-url/internal/routes"
	"go-service-shortener-url/internal/tools"
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
	tools.InfoLog.Printf("Build version: %s", buildVersion)
	tools.InfoLog.Printf("Build date: %s", buildDate)
	tools.InfoLog.Printf("Build commit: %s", buildCommit)

	config.CheckFlagEnv()
	config.InitBD()

	address := config.Env.Address

	r := routes.SetupRouter()
	pprof.Register(r)

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("myPrettyHttpsServerForYandex.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	if config.Env.EnableHTTPS != "" {
		log.Fatal(autotls.RunWithManager(r, &m))
	} else {
		if err := r.Run(address); err != nil {
			tools.ErrorLog.Fatal(err)
		}
	}
}
