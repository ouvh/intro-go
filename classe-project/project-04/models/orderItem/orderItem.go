package orderItem

import (
	"um6p.ma/project-04/models/book"
)

type OrderItem struct {
	Book     book.Book `json:"book"`
	Quantity int       `json:"quantity"`
}
