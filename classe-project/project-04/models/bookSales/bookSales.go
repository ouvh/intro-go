package bookSales

import "um6p.ma/project-04/models/book"

type BookSales struct {
	Book     book.Book `json:"book"`
	Quantity int       `json:"quantity_sold"`
}
