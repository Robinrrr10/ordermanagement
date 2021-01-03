package main

import (
	"fmt"

	"./cmd"
	"./pkg/utils"
)

func main() {
	fmt.Println("Start")
	utils.ReadProperty()
	cmd.Start()
	fmt.Println("End")
}
