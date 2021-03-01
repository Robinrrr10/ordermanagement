package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"../dbUtils"
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
			{
				orderResponse := getOrderDetail(storeId, orderId)
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusCreated)
				json.NewEncoder(res).Encode(orderResponse)
			}

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

	query := "INSERT INTO orders (storeId, productId, productName, eachPrice, quantity, totalPrice) values (" + strconv.Itoa(orderDetailRequest.StoreId) + ", " + strconv.Itoa(orderDetailRequest.ProductId) + ", '" + orderDetailRequest.ProductName + "', " + strconv.FormatFloat(orderDetailRequest.EachPrice, 'f', 6, 64) + ", " + strconv.Itoa(orderDetailRequest.Quantity) + ", " + strconv.FormatFloat(orderDetailRequest.TotalPrice, 'f', 2, 64) + ");"

	result := dbUtils.Update(query)
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error:", err)
		orderResponse = entities.OrderResponse{}
		errStatus := entities.Status{StatusCode: 2001, Message: "Failed while creating the order", Result: "ERROR"}
		orderResponse.Status = errStatus
	} else {
		fmt.Println("Inserted id is:" + strconv.FormatInt(id, 10))
		orderResponse = entities.OrderResponse{}
		status := entities.Status{StatusCode: 1001, Message: "Created successfully", Result: "SUCCESS"}
		orderDetail := orderDetailRequest
		orderDetail.OrderId = int(id)
		var orderDetails [1]entities.OrderDetail
		orderDetails[0] = orderDetail
		orderResponse.Status = status
		orderResponse.Slice = orderDetails[:]
	}
	return
}

func getOrderDetail(storeId string, orderId string) entities.OrderResponse {
	query := "SELECT id, storeId, productId, productName, eachPrice, quantity, totalPrice FROM orders WHERE storeId='" + storeId + "' AND id='" + orderId + "';"
	result := dbUtils.Fetch(query)
	orderResponse := entities.OrderResponse{}
	if result.Next() {
		orderDetail := entities.OrderDetail{}
		err := result.Scan(&orderDetail.OrderId, &orderDetail.StoreId, &orderDetail.ProductId, &orderDetail.ProductName, &orderDetail.EachPrice, &orderDetail.Quantity, &orderDetail.TotalPrice)
		if err != nil {
			status := entities.Status{StatusCode: 2002, Message: err.Error(), Result: "ERROR"}
			orderResponse.Status = status
		} else {
			//fmt.Println("Product name is:", val[3])
			//orderDetail := entities.OrderDetail{}
			//orderDetail.OrderId, _ = strconv.Atoi(val[0])
			//orderDetail.StoreId, _ = strconv.Atoi(val[1])
			//orderDetail.ProductId, _ = strconv.Atoi(val[2])
			//orderDetail.ProductName = val[3]
			//orderDetail.EachPrice, _ = strconv.ParseFloat(val[4], 64)
			//orderDetail.Quantity, _ = strconv.Atoi(val[5])
			//orderDetail.TotalPrice, _ = strconv.ParseFloat(val[6], 64)
			var orderDetails [1]entities.OrderDetail
			orderDetails[0] = orderDetail
			status := entities.Status{StatusCode: 1002, Message: "Fetched order detail successfully", Result: "SUCCESS"}
			orderResponse.Status = status
			orderResponse.Slice = orderDetails[:]
		}
	} else {
		status := entities.Status{StatusCode: 2003, Message: "Error when fetching order detail", Result: "ERROR"}
		orderResponse.Status = status
	}
	return orderResponse
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
