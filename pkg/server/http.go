package server

import (
	"fmt"
	"net/http"

	"../utils"
)

func HttpServer() {
	fmt.Println("Start Server")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hi. Hello!!!")
	})

	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "UP")
	})

	http.HandleFunc("/store/", StoreOrder)
	http.ListenAndServe(":"+utils.ServerPort, nil)
}
