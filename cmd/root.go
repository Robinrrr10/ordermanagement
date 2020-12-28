package cmd

import (
	"fmt"

	"../pkg/server"
)

func Start() {
	fmt.Println("in Root")
	server.HttpServer()
}
