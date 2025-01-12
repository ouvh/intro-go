package bookStore

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	model "um6p.ma/project-04/models/book"
	"um6p.ma/project-04/store/searchCriteria"
)

type BookExecutor interface {
	CreateBook(book model.Book) (model.Book, error)
	GetBook(id int) (model.Book, error)
	UpdateBook(id int, book model.Book) (model.Book, error)
	DeleteBook(id int) error
	SearchBooks(criteria searchCriteria.SearchCriteria) ([]model.Book, error)
}
type BookStore struct {
	sync.Mutex
	Books []model.Book
	Index int64
}

func (b *BookStore) CreateBook(book model.Book) (model.Book, error) {

	b.Lock()
	defer b.Unlock()
	b.Index++
	book.ID = (b.Index)
	book.PublishedAt = time.Now()
	b.Books = append(b.Books, book)
	return book, nil

}

func (b *BookStore) GetBook(id int) (model.Book, error) {

	b.Lock()
	defer b.Unlock()
	for _, book := range b.Books {
		if book.ID == int64(id) {
			return book, nil
		}
	}
	return model.Book{}, errors.New("book not found")

}

func (b *BookStore) UpdateBook(id int, book model.Book) (model.Book, error) {
	book.ID = int64(id)
	b.Lock()
	defer b.Unlock()
	for index, book_ := range b.Books {
		if book_.ID == int64(id) {
			book.PublishedAt = book_.PublishedAt
			b.Books[index] = book
			return book, nil
		}
	}
	return book, errors.New("book not found")

}

func (b *BookStore) DeleteBook(id int) error {
	b.Lock()
	defer b.Unlock()
	for index, book := range b.Books {
		if book.ID == int64(id) {
			b.Books = append(b.Books[:index], b.Books[index+1:]...)
			return nil
		}
	}
	return errors.New("book not found")

}

func (b *BookStore) SearchBooks(criteria searchCriteria.SearchCriteria) ([]model.Book, error) {
	b.Lock()
	defer b.Unlock()
	results := make([]model.Book, 0)

	if len(criteria.Parameters) == 0 {
		return b.Books, nil
	}

	for _, book := range b.Books {
		matched := true
	loop:
		for key, value := range criteria.Parameters {
			comp, exist := model.ComparableFields[key]
			if exist {
				if comp == 1 {
					v, err := model.GetField(book, key)
					if err != nil {
						return results, err
					}
					if reflect.ValueOf(v).Type() == reflect.ValueOf(value).Type() {
						switch reflect.ValueOf(v).Kind() {
						case reflect.Int, reflect.Int32, reflect.Int64:
							if reflect.ValueOf(v).Int() > reflect.ValueOf(value).Int() {
								matched = false
								break loop
							}
						case reflect.Float32, reflect.Float64:
							if reflect.ValueOf(v).Float() > reflect.ValueOf(value).Float() {
								matched = false
								break loop

							}
						case reflect.String:
							if reflect.ValueOf(v).String() > reflect.ValueOf(value).String() {
								matched = false
								break loop

							}

						default:
							vv := v.(time.Time)
							vvvalue := value.(time.Time)
							if vv.After(vvvalue) {
								matched = false
								break loop

							}
						}
					} else {
						return results, fmt.Errorf("type mismatch: %T vs %T", v, value)
					}
				}

			} else {
				f, err := model.GetField(book, key)
				if err != nil {
					return results, err
				}

				if value != f {
					matched = false
					break
				}

			}

		}

		if matched {
			results = append(results, book)
		}

	}

	return results, nil

}
