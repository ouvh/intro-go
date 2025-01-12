package orderStore

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	model "um6p.ma/project-04/models/order"
	"um6p.ma/project-04/store/searchCriteria"
)

type OrderExecutor interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrder(id int) (model.Order, error)
	UpdateOrder(id int, order model.Order) (model.Order, error)
	DeleteOrder(id int) error
	SearchOrder(criteria searchCriteria.SearchCriteria) ([]model.Order, error)
}

type OrderStore struct {
	sync.Mutex
	Orders []model.Order
	Index  int64
}

func (b *OrderStore) CreateOrder(order model.Order) (model.Order, error) {

	b.Lock()
	defer b.Unlock()
	b.Index++
	order.ID = int64(b.Index)
	order.CreatedAt = time.Now()
	b.Orders = append(b.Orders, order)
	return order, nil

}

func (b *OrderStore) GetOrder(id int) (model.Order, error) {

	b.Lock()
	defer b.Unlock()
	for _, order := range b.Orders {
		if order.ID == int64(id) {
			return order, nil
		}
	}
	return model.Order{}, errors.New("order not found")

}

func (b *OrderStore) UpdateOrder(id int, order model.Order) (model.Order, error) {
	order.ID = int64(id)
	b.Lock()
	defer b.Unlock()
	for index, order_ := range b.Orders {
		if order_.ID == int64(id) {
			order.CreatedAt = order_.CreatedAt
			b.Orders[index] = order
			return order, nil
		}
	}
	return order, errors.New("order not found")

}

func (b *OrderStore) DeleteOrder(id int) error {
	b.Lock()
	defer b.Unlock()
	for index, order := range b.Orders {
		if order.ID == int64(id) {
			b.Orders = append(b.Orders[:index], b.Orders[index+1:]...)
			return nil
		}
	}
	return errors.New("order not found")

}

func (b *OrderStore) SearchOrder(criteria searchCriteria.SearchCriteria) ([]model.Order, error) {
	b.Lock()
	defer b.Unlock()
	results := make([]model.Order, 0)

	if len(criteria.Parameters) == 0 {
		return b.Orders, nil
	}

	for _, order := range b.Orders {
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
							vvvalue, err := time.Parse("2006-01-02T15:04:05", value.(string))
							if err != nil {
								return results, fmt.Errorf("datetime format incompatible")
							}
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
				f, err := model.GetField(order, key)
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
			results = append(results, order)
		}

	}

	return results, nil

}
