package main

import "example.com/m/v2/routes"

func main() {
	r := routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
