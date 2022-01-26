package basket

type ShoppingCartItem struct {
	Quantity    int     `json:"quantity"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
}

type ShoppingCart struct {
	Username string             `json:"username"`
	Items    []ShoppingCartItem `json:"items"`
}

func (s *ShoppingCart) TotalPrice() float64 {
	totalPrice := 0.0
	for _, item := range s.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}
	return totalPrice
}
