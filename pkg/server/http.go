package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

	http.HandleFunc("/store/", func(res http.ResponseWriter, req *http.Request) {
		endPoint := req.URL.Path
		storeId, orderId := getStoreIdAndOrderId(endPoint)
		if req.Method == "POST" && orderId == "" && storeId != "" {
			body, _ := ioutil.ReadAll(req.Body)
			fmt.Println("body is:" + string(body))
			fmt.Fprintf(res, "create order")
		} else if req.Method == "GET" && orderId == "all" && storeId != "" {
			fmt.Fprintf(res, "get all orders")
		} else if (req.Method == "GET" || req.Method == "PUT" || req.Method == "DELETE") && orderId != "" && storeId != "" {
			switch req.Method {
			case http.MethodGet:
				fmt.Fprintf(res, "order id is there, get order detail")
			case http.MethodPut:
				fmt.Fprintf(res, "order id is there, edit order detail")
			case http.MethodDelete:
				fmt.Fprintf(res, "order id is there, delete order")
			default:
				fmt.Fprintf(res, "Invalid method")
			}
		} else {
			fmt.Fprintf(res, "Invalid request")
		}
	})

	http.ListenAndServe(":"+utils.ServerPort, nil)
}

func getStoreIdAndOrderId(endPoint string) (storeId string, orderId string) {
	stStartInx := strings.Index(endPoint, "store/")
	endStoreInx := strings.Index(endPoint, "order")
	fmt.Println("Start:", stStartInx, " End:", endStoreInx)
	storeId = ""
	storeId = endPoint[stStartInx+6 : endStoreInx-1]
	fmt.Println("StoreId::", storeId)
	endPointAfterStoreId := endPoint[endStoreInx+5:]
	orderId = ""
	if strings.Contains(endPointAfterStoreId, "/") {
		slashIndex := strings.Index(endPointAfterStoreId, "/")
		orderId = endPointAfterStoreId[slashIndex+1:]
	}
	fmt.Println("OrderId::", orderId)
	return
}
