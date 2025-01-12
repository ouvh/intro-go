package authorStore

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"

	model "um6p.ma/project-04/models/author"
	"um6p.ma/project-04/store/searchCriteria"
)

type AuthorExecutor interface {
	CreateAuthor(author model.Author) (model.Author, error)
	GetAuthor(id int) (model.Author, error)
	UpdateAuthor(id int, author model.Author) (model.Author, error)
	DeleteAuthor(id int) error
	SearchAuthor(criteria searchCriteria.SearchCriteria) ([]model.Author, error)
}

type AuthorStore struct {
	sync.RWMutex
	Authors []model.Author
	Index   int64
}

func (b *AuthorStore) CreateAuthor(author model.Author) (model.Author, error) {
	b.Lock()
	defer b.Unlock()
	b.Index++
	author.ID = int64(b.Index)
	b.Authors = append(b.Authors, author)

	return author, nil

}

func (b *AuthorStore) GetAuthor(id int) (model.Author, error) {

	b.Lock()
	defer b.Unlock()
	for _, author := range b.Authors {
		if author.ID == int64(id) {
			return author, nil
		}
	}
	return model.Author{}, errors.New("author not found")

}

func (b *AuthorStore) UpdateAuthor(id int, author model.Author) (model.Author, error) {
	author.ID = int64(id)
	b.Lock()
	defer b.Unlock()
	for index, author_ := range b.Authors {
		if author_.ID == int64(id) {
			b.Authors[index] = author
			return author, nil
		}
	}
	return author, errors.New("author not found")

}

func (b *AuthorStore) DeleteAuthor(id int) error {
	b.Lock()
	defer b.Unlock()
	for index, author := range b.Authors {
		if author.ID == int64(id) {
			b.Authors = append(b.Authors[:index], b.Authors[index+1:]...)
			return nil
		}
	}
	return errors.New("author not found")

}

func (b *AuthorStore) SearchAuthor(criteria searchCriteria.SearchCriteria) ([]model.Author, error) {
	b.Lock()
	defer b.Unlock()
	results := make([]model.Author, 0)

	if len(criteria.Parameters) == 0 {
		return b.Authors, nil
	}

	for _, order := range b.Authors {
		matched := true
	loop:
		for key, value := range criteria.Parameters {
			comp, exist := model.ComparableFields[key]
			if exist {
				if comp == 1 {
					v, err := model.GetField(order, key)
					if err != nil {
						return results, err
					}
					if true {
						x, err := strconv.ParseFloat(value.(string), 64)
						if err != nil {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if reflect.ValueOf(v).Float() > reflect.ValueOf(value).Float() {
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
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.After(vvvalue) {
									matched = false
									break loop

								}

							}

						} else {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if (reflect.ValueOf(v).Float)() > reflect.ValueOf(x).Float() {
									matched = false
									break loop
								}
							case reflect.Float32, reflect.Float64:
								if reflect.ValueOf(v).Float() > reflect.ValueOf(x).Float() {
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
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.After(vvvalue) {
									matched = false
									break loop

								}

							}

						}

					} else {
						return results, fmt.Errorf("type mismatch: %T vs %T", v, value)
					}
				} else {

					v, err := model.GetField(order, key)
					if err != nil {
						return results, err
					}
					if true {
						x, err := strconv.ParseFloat(value.(string), 64)
						if err != nil {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if reflect.ValueOf(v).Float() < reflect.ValueOf(value).Float() {
									matched = false
									break loop
								}
							case reflect.Float32, reflect.Float64:
								if reflect.ValueOf(v).Float() < reflect.ValueOf(value).Float() {
									matched = false
									break loop

								}
							case reflect.String:
								if reflect.ValueOf(v).String() < reflect.ValueOf(value).String() {
									matched = false
									break loop

								}

							default:
								vv := v.(time.Time)
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.Before(vvvalue) {
									matched = false
									break loop

								}

							}

						} else {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if (reflect.ValueOf(v).Float)() < reflect.ValueOf(x).Float() {
									matched = false
									break loop
								}
							case reflect.Float32, reflect.Float64:
								if reflect.ValueOf(v).Float() < reflect.ValueOf(x).Float() {
									matched = false
									break loop

								}
							case reflect.String:
								if reflect.ValueOf(v).String() < reflect.ValueOf(value).String() {
									matched = false
									break loop

								}

							default:
								vv := v.(time.Time)
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.Before(vvvalue) {
									matched = false
									break loop

								}

							}

						}

					} else {
						return results, fmt.Errorf("type mismatch: %T vs %T", v, value)
					}

				}

			} else {
				f, err := model.GetField(order, key)
				if err != nil {
					return results, err
				}

				x, err := strconv.ParseFloat(value.(string), 64)
				if err != nil {
					if value != f {
						matched = false
						break
					}
				} else {
					if float64(x) != float64(f.(float64)) {
						matched = false
						break
					}
				}

			}

		}

		if matched {
			results = append(results, order)
		}

	}

	return results, nil
}
