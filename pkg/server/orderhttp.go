package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"../entities"
)

func StoreOrder(res http.ResponseWriter, req *http.Request) {
	endPoint := req.URL.Path
	storeId, orderId := getStoreIdAndOrderId(endPoint)
	if req.Method == "POST" && orderId == "" && storeId != "" {
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println("body is:" + string(body))
		var orderDetail entities.OrderDetail
		json.Unmarshal(body, &orderDetail)
		//fmt.Fprintf(res, "create order. ProductName: %v", orderDetail.ProductName)
		orderResponse := createOrder(storeId, orderDetail)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(orderResponse)
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
}

func createOrder(storeId string, orderDetailRequest entities.OrderDetail) (orderResponse entities.OrderResponse) {

	fmt.Println("Product is:" + orderDetailRequest.ProductName)

	//Call db to insert the value

	orderResponse = entities.OrderResponse{}
	status := entities.Status{StatusCode: 1001, Message: "Created successfully", Result: "SUCCESS"}
	orderDetail := entities.OrderDetail{OrderId: 1001, ProductName: "Bullet belt"}
	var orderDetails [1]entities.OrderDetail
	orderDetails[0] = orderDetail
	orderResponse.Status = status
	orderResponse.Slice = orderDetails[:]
	return
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
