package customerService

import (
	"um6p.ma/project-04/dao"
	model "um6p.ma/project-04/models/customer"
	"um6p.ma/project-04/store"
	"um6p.ma/project-04/store/searchCriteria"
)

type CustomerServiceFunctionalities interface {
	CreateCustomer(Customer model.Customer) (model.Customer, error)
	GetCustomer(id int) (model.Customer, error)
	UpdateCustomer(id int, Customer model.Customer) (model.Customer, error)
	DeleteCustomer(id int) error
	SearchCustomer(criteria searchCriteria.SearchCriteria) ([]model.Customer, error)
}

type CustomerService struct {
	Store *store.Store
}

func (b *CustomerService) CreateCustomer(customer dao.CustomerDAO) (model.Customer, error) {

	cu := model.Customer{ID: 0, Name: customer.Name, Email: customer.Email, Address: model.Address{Street: customer.Address.Street, City: customer.Address.City, State: customer.Address.State, PostalCode: customer.Address.PostalCode, Country: customer.Address.Country}}
	result, err := b.Store.CustomerStore.CreateCustomer(cu)
	return result, err

}

func (b *CustomerService) GetCustomer(id int) (model.Customer, error) {

	result, err := b.Store.CustomerStore.GetCustomer(id)
	return result, err

}

func (b *CustomerService) UpdateCustomer(id int, customer dao.CustomerDAO) (model.Customer, error) {

	cu := model.Customer{ID: 0, Name: customer.Name, Email: customer.Email, Address: model.Address{Street: customer.Address.Street, City: customer.Address.City, State: customer.Address.State, PostalCode: customer.Address.PostalCode, Country: customer.Address.Country}}
	result, err := b.Store.CustomerStore.UpdateCustomer(id, cu)
	return result, err

}

func (b *CustomerService) DeleteCustomer(id int) error {
	err := b.Store.CustomerStore.DeleteCustomer(id)
	return err

}

func (b *CustomerService) SearchCustomer(criteria searchCriteria.SearchCriteria) ([]model.Customer, error) {
	results, err := b.Store.CustomerStore.SearchCustomer(criteria)
	return results, err

}
