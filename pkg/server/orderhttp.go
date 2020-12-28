package server

import (
	"fmt"
	"net/http"
	"strings"

	"../dbUtils"
)

func StoreOrder(res http.ResponseWriter, req *http.Request) {
	endPoint := req.URL.Path
	stStartInx := strings.Index(endPoint, "store/")
	endStoreInx := strings.Index(endPoint, "order/")
	fmt.Println("Start:", stStartInx, " End:", endStoreInx)
	storeId := endPoint[stStartInx+6 : endStoreInx-1]
	orderId := endPoint[endStoreInx+6:]
	fmt.Println("StoreId::", storeId, " OrderId::", orderId)
	dbUtils.Connect()
	fmt.Fprintf(res, "working...Check console logs")
}
