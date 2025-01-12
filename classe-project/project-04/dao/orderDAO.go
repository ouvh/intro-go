package dao

type OrderItemDAO struct {
	BookId   int64 `json:"book"`
	Quantity int   `json:"quantity"`
}

type OrderDAO struct {
	CustomerId int64          `json:"customer_id"`
	Items      []OrderItemDAO `json:"items"`
	Status     string         `json:"status"`
}
