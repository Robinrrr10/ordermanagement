package entities

type OrderDetail struct {
	OrderId     int     `json:"orderId"`
	StoreId     int     `json:"storeId"`
	ProductId   int     `json:"productId"`
	ProductName string  `json:"productName"`
	EachPrice   float64 `json:"eachPrice"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"totalPrice"`
}

type Status struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Result     string `json:"result"`
}

type OrderResponse struct {
	Status Status        `json:"status"`
	Slice  []OrderDetail `json:"data"`
}
