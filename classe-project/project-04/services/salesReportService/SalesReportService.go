package SalesReportService

import (
	"sync"

	"um6p.ma/project-04/models/bookSales"
	model "um6p.ma/project-04/models/salesReport"
	"um6p.ma/project-04/store"
	"um6p.ma/project-04/store/searchCriteria"
)

type SalesReportFunctionalities interface {
	CreateSalesReport() (model.SalesReport, error)
	SearchSalesReport(criteria searchCriteria.SearchCriteria) ([]model.SalesReport, error)
}

type SalesReportService struct {
	sync.Mutex
	Store *store.Store
}

func (b *SalesReportService) CreateSalesReport() error {
	sales := b.Store.SalesReportStore.RetrieveSales()

	if len(sales) == 0 {
		b.Store.SalesReportStore.CreateSalesReport(model.SalesReport{TotalRevenue: 0, TotalOrders: 0, TopSellingBooks: make([]bookSales.BookSales, 0)})
	}

	booksales := make(map[int64]bookSales.BookSales, 0)
	totalRevenue := float64(0)
	TotalOrders := len(sales)
	for _, sale := range sales {
		totalRevenue = totalRevenue + sale.TotalPrice
		for _, item := range sale.Items {
			_, exist := booksales[item.Book.ID]
			if exist {
				temp := booksales[item.Book.ID]
				temp.Book = item.Book
				temp.Quantity = temp.Quantity + item.Quantity
				booksales[item.Book.ID] = temp

			} else {
				temp := bookSales.BookSales{Quantity: 0}
				temp.Book = item.Book
				temp.Quantity = temp.Quantity + item.Quantity
				booksales[item.Book.ID] = temp
			}
		}
	}
	r := make([]bookSales.BookSales, 0)
	top := bookSales.BookSales{Quantity: 0}
	for _, value := range booksales {
		if value.Quantity > top.Quantity {
			top = value
		}
	}
	for _, value := range booksales {
		if value.Quantity == top.Quantity {
			r = append(r, value)
		}
	}
	result := model.SalesReport{TotalRevenue: totalRevenue, TotalOrders: TotalOrders, TopSellingBooks: r}
	b.Store.SalesReportStore.CreateSalesReport(result)

	return nil
}

func (b *SalesReportService) SearchSalesReport(criteria searchCriteria.SearchCriteria) ([]model.SalesReport, error) {
	results, err := b.Store.SalesReportStore.SearchSalesReport(criteria)
	return results, err
}
