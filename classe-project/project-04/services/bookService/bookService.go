package bookService

import (
	"um6p.ma/project-04/dao"
	model "um6p.ma/project-04/models/book"
	"um6p.ma/project-04/store"
	"um6p.ma/project-04/store/searchCriteria"
)

type BookServiceFunctionalities interface {
	CreateBook(Book model.Book) (model.Book, error)
	GetBook(id int) (model.Book, error)
	UpdateBook(id int, Book model.Book) (model.Book, error)
	DeleteBook(id int) error
	SearchBook(criteria searchCriteria.SearchCriteria) ([]model.Book, error)
}

type BookService struct {
	Store *store.Store
}

func (b *BookService) CreateBook(book dao.BookDAO) (model.Book, error) {
	author, err := b.Store.AuthorStore.GetAuthor(int(book.AuthorId))
	if err != nil {
		return model.Book{}, err
	}
	bo := model.Book{ID: 0, Title: book.Title, Author: author, Genres: book.Genres, Price: book.Price, Stock: book.Stock}
	result, err := b.Store.BookStore.CreateBook(bo)
	return result, err

}

func (b *BookService) GetBook(id int) (model.Book, error) {

	result, err := b.Store.BookStore.GetBook(id)
	return result, err

}

func (b *BookService) UpdateBook(id int, book dao.BookDAO) (model.Book, error) {
	author, err := b.Store.AuthorStore.GetAuthor(int(book.AuthorId))
	if err != nil {
		return model.Book{}, err
	}
	bo := model.Book{ID: 0, Title: book.Title, Author: author, Genres: book.Genres, Price: book.Price, Stock: book.Stock}
	result, err := b.Store.BookStore.CreateBook(bo)
	return result, err

}

func (b *BookService) DeleteBook(id int) error {
	err := b.Store.BookStore.DeleteBook(id)
	return err

}

func (b *BookService) SearchBook(criteria searchCriteria.SearchCriteria) ([]model.Book, error) {
	results, err := b.Store.BookStore.SearchBooks(criteria)
	return results, err

}
