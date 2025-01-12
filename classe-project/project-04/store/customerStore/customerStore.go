package customerStore

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	model "um6p.ma/project-04/models/customer"
	"um6p.ma/project-04/store/searchCriteria"
)

type CustomerExecutor interface {
	Createcustomer(customer model.Customer) (model.Customer, error)
	Getcustomer(id int) (model.Customer, error)
	Updatecustomer(id int, customer model.Customer) (model.Customer, error)
	Deletecustomer(id int) error
	Searchcustomer(criteria searchCriteria.SearchCriteria) ([]model.Customer, error)
}

type CustomerStore struct {
	sync.Mutex
	Customers []model.Customer
	Index     int64
}

func (b *CustomerStore) CreateCustomer(customer model.Customer) (model.Customer, error) {

	b.Lock()
	defer b.Unlock()
	b.Index++
	customer.ID = int64(b.Index)
	customer.CreatedAt = time.Now()
	b.Customers = append(b.Customers, customer)
	return customer, nil

}

func (b *CustomerStore) GetCustomer(id int) (model.Customer, error) {

	b.Lock()
	defer b.Unlock()
	for _, customer := range b.Customers {
		if customer.ID == int64(id) {
			return customer, nil
		}
	}
	return model.Customer{}, errors.New("customer not found")

}

func (b *CustomerStore) UpdateCustomer(id int, customer model.Customer) (model.Customer, error) {
	customer.ID = int64(id)
	b.Lock()
	defer b.Unlock()
	for index, customer_ := range b.Customers {
		if customer_.ID == int64(id) {
			customer.CreatedAt = customer_.CreatedAt
			b.Customers[index] = customer
			return customer, nil
		}
	}
	return customer, errors.New("customer not found")

}

func (b *CustomerStore) DeleteCustomer(id int) error {
	b.Lock()
	defer b.Unlock()
	for index, customer := range b.Customers {
		if customer.ID == int64(id) {
			b.Customers = append(b.Customers[:index], b.Customers[index+1:]...)
			return nil
		}
	}
	return errors.New("customer not found")

}

func (b *CustomerStore) SearchCustomer(criteria searchCriteria.SearchCriteria) ([]model.Customer, error) {
	b.Lock()
	b.Lock()
	defer b.Unlock()
	results := make([]model.Customer, 0)

	if len(criteria.Parameters) == 0 {
		return b.Customers, nil
	}

	for _, customer := range b.Customers {
		matched := true
	loop:
		for key, value := range criteria.Parameters {
			comp, exist := model.ComparableFields[key]
			if exist {
				if comp == 1 {
					v, err := model.GetField(customer, key)
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
				f, err := model.GetField(customer, key)
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
			results = append(results, customer)
		}

	}

	return results, nil

}
