package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Robinrrr10/ordermanagement/pkg/dbUtils"
	"github.com/Robinrrr10/ordermanagement/pkg/entities"
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
		orderResponse := getAllOrders(storeId)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(orderResponse)
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
			{
				body, _ := ioutil.ReadAll(req.Body)
				fmt.Println("body is:" + string(body))
				var orderDetail entities.OrderDetail
				json.Unmarshal(body, &orderDetail)
				orderResponse := modifyOrder(storeId, orderId, orderDetail)
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusCreated)
				json.NewEncoder(res).Encode(orderResponse)
			}

		case http.MethodDelete:
			{
				orderResponse := removeOrder(storeId, orderId)
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusCreated)
				json.NewEncoder(res).Encode(orderResponse)
			}
		default:
			{
				orderResponse := entities.OrderResponse{}
				errStatus := entities.Status{StatusCode: 2001, Message: "Invalid request", Result: "ERROR"}
				orderResponse.Status = errStatus
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusCreated)
				json.NewEncoder(res).Encode(orderResponse)
			}
		}
	} else {
		orderResponse := entities.OrderResponse{}
		errStatus := entities.Status{StatusCode: 2001, Message: "Invalid request", Result: "ERROR"}
		orderResponse.Status = errStatus
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(orderResponse)
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

func modifyOrder(storeId string, orderId string, orderDetailRequest entities.OrderDetail) (orderResponse entities.OrderResponse) {

	availableOrderDetail := getOrderDetail(storeId, orderId)

	if orderDetailRequest.ProductId == 0 {
		orderDetailRequest.ProductId = availableOrderDetail.Slice[0].ProductId
	}
	if orderDetailRequest.ProductName == "" {
		orderDetailRequest.ProductName = availableOrderDetail.Slice[0].ProductName
	}
	if orderDetailRequest.EachPrice == 0 {
		orderDetailRequest.EachPrice = availableOrderDetail.Slice[0].EachPrice
	}
	if orderDetailRequest.Quantity == 0 {
		orderDetailRequest.Quantity = availableOrderDetail.Slice[0].Quantity
	}
	if orderDetailRequest.TotalPrice == 0 {
		orderDetailRequest.TotalPrice = availableOrderDetail.Slice[0].TotalPrice
	}

	query := "UPDATE orders SET productId=" + strconv.Itoa(orderDetailRequest.ProductId) + ", productName='" + orderDetailRequest.ProductName + "', eachPrice=" + strconv.FormatFloat(orderDetailRequest.EachPrice, 'f', 6, 64) + ", quantity=" + strconv.Itoa(orderDetailRequest.Quantity) + ", totalPrice=" + strconv.FormatFloat(orderDetailRequest.TotalPrice, 'f', 2, 64) + " WHERE storeId='" + storeId + "' AND id='" + orderId + "';"

	result := dbUtils.Update(query)
	id, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error:", err)
		orderResponse = entities.OrderResponse{}
		errStatus := entities.Status{StatusCode: 2001, Message: "Failed while creating the order", Result: "ERROR"}
		orderResponse.Status = errStatus
	} else {
		fmt.Println("updated number of orders is:" + strconv.FormatInt(id, 10))
		orderResponse = entities.OrderResponse{}
		status := entities.Status{StatusCode: 1001, Message: "Order updated successfully", Result: "SUCCESS"}
		orderDetail := orderDetailRequest
		orderDetail.OrderId = availableOrderDetail.Slice[0].OrderId
		orderDetail.StoreId = availableOrderDetail.Slice[0].StoreId
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

func removeOrder(storeId string, orderId string) (orderResponse entities.OrderResponse) {
	query := "DELETE FROM orders WHERE storeId='" + storeId + "' AND orderId='" + orderId + "';"
	result := dbUtils.Update(query)
	rowsAff, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error:", err)
		orderResponse = entities.OrderResponse{}
		errStatus := entities.Status{StatusCode: 2001, Message: "Failed while deleting the order", Result: "ERROR"}
		orderResponse.Status = errStatus
	} else if rowsAff == 1 {
		fmt.Println("number of deleted records:" + strconv.FormatInt(rowsAff, 10))
		orderResponse = entities.OrderResponse{}
		status := entities.Status{StatusCode: 1001, Message: "Deleted successfully", Result: "SUCCESS"}
		orderResponse.Status = status
	} else {
		fmt.Println("number of deleted records:" + strconv.FormatInt(rowsAff, 10))
		orderResponse = entities.OrderResponse{}
		status := entities.Status{StatusCode: 1001, Message: "Deleting Unsuccessfully", Result: "SUCCESS"}
		orderResponse.Status = status
	}
	return
}

func getAllOrders(storeId string) entities.OrderResponse {
	query := "SELECT id, storeId, productId, productName, eachPrice, quantity, totalPrice FROM orders WHERE storeId='" + storeId + "';"
	result := dbUtils.Fetch(query)
	orderResponse := entities.OrderResponse{}
	var isOrderAvailable bool = false
	var orderDetails [0]entities.OrderDetail
	orderResponse.Slice = orderDetails[:]

	for result.Next() {
		orderDetail := entities.OrderDetail{}
		err := result.Scan(&orderDetail.OrderId, &orderDetail.StoreId, &orderDetail.ProductId, &orderDetail.ProductName, &orderDetail.EachPrice, &orderDetail.Quantity, &orderDetail.TotalPrice)
		if err != nil {
			break
		} else {
			isOrderAvailable = true
			orderResponse.Slice = append(orderResponse.Slice, orderDetail)
		}
	}

	if isOrderAvailable {
		status := entities.Status{StatusCode: 1002, Message: "Fetched order detail successfully", Result: "SUCCESS"}
		orderResponse.Status = status
	} else {
		status := entities.Status{StatusCode: 2003, Message: "No order available", Result: "ERROR"}
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
