package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"example.com/m/v2/internal/config"
	"example.com/m/v2/internal/routes"
	"github.com/gin-contrib/pprof"
)

var (
	buildDate    string
	buildVersion string
	buildCommit  string
)

// main основная точка входа приложения
func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	defer stop()

	println(fmt.Sprintf("Build version: %s", buildVersion))
	println(fmt.Sprintf("Build date: %s", buildDate))
	println(fmt.Sprintf("Build commit: %s", buildCommit))

	config.CheckFlagEnv()
	config.InitBD()

	address := config.Env.Address
	pem := config.Env.Pem
	key := config.Env.Key

	r := routes.SetupRouter()
	pprof.Register(r)

	go func() {
		if config.Env.EnableHTTPS {
			err := r.RunTLS(address, pem, key)
			if err != nil {
				panic(err)
			}
		} else {
			if err := r.Run(address); err != nil {
				panic(err)
			}
		}
	}()

	<-ctx.Done()
	if ctx.Err() != nil {
		fmt.Printf("Приложение завершенно сигналом: %v\n", ctx.Err())
	}
}
