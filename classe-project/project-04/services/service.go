package service

import (
	"time"

	"um6p.ma/project-04/services/authorService"
	"um6p.ma/project-04/services/bookService"
	"um6p.ma/project-04/services/customerService"
	"um6p.ma/project-04/services/orderService"
	SalesReportService "um6p.ma/project-04/services/salesReportService"

	"um6p.ma/project-04/store"
)

type Service struct {
	BookService        bookService.BookService
	CustomerService    customerService.CustomerService
	AuthorService      authorService.AuthorService
	OrderService       orderService.OrderService
	SalesReportService SalesReportService.SalesReportService
	Storage            *store.Store
	ReportDuration     time.Duration
}

func (s *Service) Init(store *store.Store) {
	// Dependency injection
	s.Storage = store
	s.BookService.Store = store
	s.CustomerService.Store = store
	s.AuthorService.Store = store
	s.OrderService.Store = store
	s.SalesReportService.Store = store

	go func() {
		for {
			time.Sleep(s.ReportDuration)
			s.SalesReportService.CreateSalesReport()
		}

	}()
}
