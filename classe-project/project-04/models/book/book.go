package book

import (
	"fmt"
	"time"

	"um6p.ma/project-04/models/author"
)

type Book struct {
	ID          int64         `json:"id"`
	Title       string        `json:"title"`
	Author      author.Author `json:"author"`
	Genres      []string      `json:"genres"`
	PublishedAt time.Time     `json:"published_at"`
	Price       float64       `json:"price"`
	Stock       int           `json:"stock"`
}

var Fields = []string{"ID", "Title", "AuthorID", "AuthorFirstName", "AuthorLastName", "AuthorBio", "PublishedAt", "Price", "Stock", "PublishedAtFrom", "PublishedAtTo", "PriceFrom", "PriceTo", "StockFrom", "StockTo"}

var ComparableFields = map[string]int{"PublishedAtFrom": 0, "PublishedAtTo": 1, "PriceFrom": 0, "PriceTo": 1, "StockFrom": 0, "StockTo": 1}

func GetField(book Book, field string) (interface{}, error) {
	switch field {
	case "PublishedAtFrom":
		return book.PublishedAt, nil
	case "PublishedAtTo":
		return book.PublishedAt, nil
	case "PriceFrom":
		return book.Price, nil
	case "PriceTo":
		return book.Price, nil
	case "StockFrom":
		return book.Stock, nil
	case "StockTo":
		return book.Stock, nil
	case "ID":
		return book.ID, nil
	case "Title":
		return book.Title, nil
	case "AuthorID":
		return book.Author.ID, nil
	case "AuthorFirstName":
		return book.Author.FirstName, nil
	case "AuthorLastName":
		return book.Author.FirstName, nil
	case "AuthorBio":
		return book.Author.Bio, nil
	case "PublishedAt":
		return book.PublishedAt, nil
	case "Price":
		return book.Price, nil
	case "Stock":
		return book.Stock, nil
	default:
		return nil, fmt.Errorf("field '%s' does not exist", field)
	}
}
