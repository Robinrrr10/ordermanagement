package main

import (
	"fmt"

	"github.com/Robinrrr10/ordermanagement/cmd"
	"github.com/Robinrrr10/ordermanagement/pkg/utils"
)

func main() {
	fmt.Println("Start")
	utils.ReadProperty()
	cmd.Start()
	fmt.Println("End")
}
