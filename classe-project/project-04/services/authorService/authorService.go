package authorService

import (
	"um6p.ma/project-04/dao"
	model "um6p.ma/project-04/models/author"
	"um6p.ma/project-04/store"
	"um6p.ma/project-04/store/searchCriteria"
)

type AuthorServiceFunctionalities interface {
	CreateAuthor(author model.Author) (model.Author, error)
	GetAuthor(id int) (model.Author, error)
	UpdateAuthor(id int, author model.Author) (model.Author, error)
	DeleteAuthor(id int) error
	SearchAuthor(criteria searchCriteria.SearchCriteria) ([]model.Author, error)
}

type AuthorService struct {
	Store *store.Store
}

func (b *AuthorService) CreateAuthor(author dao.AuthorDAO) (model.Author, error) {
	au := model.Author{ID: 0, FirstName: author.FirstName, LastName: author.LastName, Bio: author.Bio}
	result, err := b.Store.AuthorStore.CreateAuthor(au)

	return result, err

}

func (b *AuthorService) GetAuthor(id int) (model.Author, error) {

	result, err := b.Store.AuthorStore.GetAuthor(id)
	return result, err

}

func (b *AuthorService) UpdateAuthor(id int, author dao.AuthorDAO) (model.Author, error) {
	au := model.Author{ID: 0, FirstName: author.FirstName, LastName: author.LastName, Bio: author.Bio}
	result, err := b.Store.AuthorStore.UpdateAuthor(id, au)
	return result, err

}

func (b *AuthorService) DeleteAuthor(id int) error {
	err := b.Store.AuthorStore.DeleteAuthor(id)
	return err

}

func (b *AuthorService) SearchAuthor(criteria searchCriteria.SearchCriteria) ([]model.Author, error) {
	results, err := b.Store.AuthorStore.SearchAuthor(criteria)
	return results, err

}
