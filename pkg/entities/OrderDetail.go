package entities

type OrderDetail struct {
	orderId            int
	productId          int
	productName        string
	eachPrice          float32
	quantity           int
	discountPercentage float32
	totalPrice         float32
}
