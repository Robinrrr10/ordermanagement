package server

import (
	"fmt"
	"net/http"
)

func HttpServer() {
	fmt.Println("Start Server")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hi. Hello!!!")
	})

	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "UP")
	})

	http.ListenAndServe(":7878", nil)
}
