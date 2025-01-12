package order

import (
	"fmt"
	"time"

	"um6p.ma/project-04/models/customer"
	"um6p.ma/project-04/models/orderItem"
)

type Order struct {
	ID         int64                 `json:"id"`
	Customer   customer.Customer     `json:"customer"`
	Items      []orderItem.OrderItem `json:"items"`
	TotalPrice float64               `json:"total_price"`
	CreatedAt  time.Time             `json:"created_at"`
	Status     string                `json:"status"`
}

var Fields = []string{"ID", "TotalPrice", "CreatedAt", "Status", "CreatedAtFrom", "CreatedAtTo", "TotalPriceFrom", "TotalPriceTo"}

var ComparableFields = map[string]int{"CreatedAtFrom": 0, "CreatedAtTo": 1, "TotalPriceFrom": 0, "TotalPriceTo": 1}

func GetField(order Order, field string) (interface{}, error) {
	switch field {
	case "CreatedAtFrom":
		return order.CreatedAt, nil
	case "CreatedAtTo":
		return order.CreatedAt, nil
	case "TotalPriceFrom":
		return order.TotalPrice, nil
	case "TotalPriceTo":
		return order.TotalPrice, nil
	case "ID":
		return order.ID, nil
	case "TotalPrice":
		return order.TotalPrice, nil
	case "CreatedAt":
		return order.CreatedAt, nil
	case "Status":
		return order.Status, nil

	default:
		return nil, fmt.Errorf("field '%s' does not exist", field)
	}
}
