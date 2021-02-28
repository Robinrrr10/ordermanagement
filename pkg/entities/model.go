package entities

type OrderDetail struct {
	OrderId            int     `json:"orderId"`
	ProductId          int     `json:"productId"`
	ProductName        string  `json:"productName"`
	EachPrice          float32 `json:"eachPrice"`
	Quantity           int     `json:"quantity"`
	DiscountPercentage float32 `json:"discountPercentage"`
	TotalPrice         float32 `json:"totalPrice"`
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
