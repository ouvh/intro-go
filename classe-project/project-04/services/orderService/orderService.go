package orderService

import (
	"fmt"
	"sync"

	"um6p.ma/project-04/dao"
	"um6p.ma/project-04/models/book"
	model "um6p.ma/project-04/models/order"
	"um6p.ma/project-04/models/orderItem"
	"um6p.ma/project-04/store"
	"um6p.ma/project-04/store/searchCriteria"
)

type OrderServiceFunctionalities interface {
	CreateOrder(Order model.Order) (model.Order, error)
	GetOrder(id int) (model.Order, error)
	UpdateOrder(id int, Order model.Order) (model.Order, error)
	DeleteOrder(id int) error
	SearchOrder(criteria searchCriteria.SearchCriteria) ([]model.Order, error)
}

type OrderService struct {
	sync.Mutex
	Store *store.Store
}

func (b *OrderService) CreateOrder(order dao.OrderDAO) (model.Order, error) {
	// finding the customer
	customer, err := b.Store.CustomerStore.GetCustomer(int(order.CustomerId))
	if err != nil {
		return model.Order{}, err
	}
	b.Lock()
	defer b.Unlock()
	// finding all the items
	history := make(map[int64]book.Book, 0)
	prices := make(map[int64]float64, 0)
	quantity := make(map[int64]int, 0)
	for _, item := range order.Items {
		book, err := b.Store.BookStore.GetBook(int(item.BookId))
		if err != nil {
			return model.Order{}, err
		}
		_, exist := history[item.BookId]
		if !exist {
			history[item.BookId] = book
			prices[item.BookId] = float64(0)
			quantity[item.BookId] = int(0)
		}
		if history[item.BookId].Stock >= item.Quantity {
			temp := history[item.BookId]
			temp.Stock = history[item.BookId].Stock - item.Quantity

			history[item.BookId] = temp
			prices[item.BookId] = prices[item.BookId] + float64(item.Quantity)*book.Price
			quantity[item.BookId] = quantity[item.BookId] + item.Quantity

		} else {
			return model.Order{}, fmt.Errorf("the Quantity Purchased for the bookid %d is not available", item.BookId)
		}

	}
	// apply changes
	items := make([]orderItem.OrderItem, 0)
	TotalPrice := float64(0)
	for key, value := range history {
		items = append(items, orderItem.OrderItem{Book: value, Quantity: quantity[key]})
		TotalPrice += prices[key]
		b.Store.BookStore.UpdateBook(int(key), value)

	}

	cu := model.Order{ID: 0, Customer: customer, Items: items, TotalPrice: TotalPrice, Status: order.Status}
	result, err := b.Store.OrderStore.CreateOrder(cu)
	// inject the new sales in the current sales list in the sales store
	b.Store.SalesReportStore.AddSale(result)
	return result, err

}

func (b *OrderService) GetOrder(id int) (model.Order, error) {

	result, err := b.Store.OrderStore.GetOrder(id)
	return result, err

}

func (b *OrderService) UpdateOrder(id int, order dao.OrderDAO) (model.Order, error) {
	return model.Order{}, fmt.Errorf("function Not Supported")
}

func (b *OrderService) DeleteOrder(id int) error {
	err := b.Store.OrderStore.DeleteOrder(id)
	return err

}

func (b *OrderService) SearchOrder(criteria searchCriteria.SearchCriteria) ([]model.Order, error) {
	results, err := b.Store.OrderStore.SearchOrder(criteria)
	return results, err

}
