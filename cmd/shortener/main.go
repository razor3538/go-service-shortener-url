package main

import (
	"example.com/m/v2/internal/config"
	pb "example.com/m/v2/internal/proto"
	"example.com/m/v2/internal/routes"
	"example.com/m/v2/internal/server"
	"fmt"
	"github.com/gin-contrib/pprof"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	buildDate    string
	buildVersion string
	buildCommit  string
)

// main основная точка входа приложения
func main() {
	//ctx, stop := signal.NotifyContext(context.Background(),
	//	syscall.SIGINT,
	//	syscall.SIGTERM,
	//	syscall.SIGQUIT,
	//	os.Interrupt)
	//
	//defer stop()

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
		if config.Env.EnableGRPC {
			listen, err := net.Listen("tcp", address)
			if err != nil {
				log.Fatal(err)
			}
			s := grpc.NewServer()

			pb.RegisterURLsServer(s, &server.UrlServer{})

			fmt.Println("Сервер gRPC начал работу")
			if err := s.Serve(listen); err != nil {
				log.Fatal(err)
			}
		}

		if config.Env.EnableHTTPS {
			err := r.RunTLS(address, pem, key)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Сервер начал работу")
			fmt.Println(address)

			if err := r.Run(address); err != nil {
				panic(err)
			}
		}
	}()

	//<-ctx.Done()
	//if ctx.Err() != nil {
	//	fmt.Printf("Приложение завершенно сигналом: %v\n", ctx.Err())
	//}
}
