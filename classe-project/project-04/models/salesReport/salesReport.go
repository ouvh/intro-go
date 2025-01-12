package salesReport

import (
	"fmt"
	"time"

	"um6p.ma/project-04/models/bookSales"
)

type SalesReport struct {
	Timestamp       time.Time             `json:"timestamp"`
	TotalRevenue    float64               `json:"total_revenue"`
	TotalOrders     int                   `json:"total_orders"`
	TopSellingBooks []bookSales.BookSales `json:"top_selling_books"`
}

var Fields = []string{"Timestamp", "TotalRevenue", "TotalOrders", "TimestampFrom", "TimestampTo", "TotalRevenueFrom", "TotalRevenueTo", "TotalOrdersFrom", "TotalOrdersTo"}
var ComparableFields = map[string]int{"TimestampFrom": 0, "TimestampTo": 1, "TotalRevenueFrom": 0, "TotalRevenueTo": 1, "TotalOrdersFrom": 0, "TotalOrdersTo": 1}

func GetField(salesReport SalesReport, field string) (interface{}, error) {
	switch field {
	case "TimestampFrom":
		return salesReport.Timestamp, nil
	case "TimestampTo":
		return salesReport.Timestamp, nil
	case "TotalRevenueFrom":
		return salesReport.TotalRevenue, nil
	case "TotalRevenueTo":
		return salesReport.Timestamp, nil
	case "TotalOrdersFrom":
		return salesReport.TotalOrders, nil
	case "TotalOrdersTo":
		return salesReport.Timestamp, nil

	case "Timestamp":
		return salesReport.Timestamp, nil
	case "TotalRevenue":
		return salesReport.TotalRevenue, nil
	case "TotalOrders":
		return salesReport.TotalOrders, nil
	default:
		return nil, fmt.Errorf("field '%s' does not exist", field)
	}
}
