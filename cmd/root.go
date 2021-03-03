package cmd

import (
	"fmt"

	"github.com/Robinrrr10/ordermanagement/pkg/server"
)

func Start() {
	fmt.Println("in Root")
	server.HttpServer()
}
