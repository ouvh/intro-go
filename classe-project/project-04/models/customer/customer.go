package customer

import (
	"fmt"
	"time"
)

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type Customer struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

var Fields = []string{"ID", "CreatedAtFrom", "CreatedAtTo", "Name", "Email", "AdressStreet", "AdressCity", "AdressState", "AdressPostalCode", "AdressCountry", "CreatedAt"}

var ComparableFields = map[string]int{"CreatedAtFrom": 0, "CreatedAtTo": 1}

func GetField(customer Customer, field string) (interface{}, error) {
	switch field {
	case "CreatedAtFrom":
		return customer.CreatedAt, nil
	case "CreatedAtTo":
		return customer.CreatedAt, nil
	case "ID":
		return customer.ID, nil
	case "Name":
		return customer.Name, nil
	case "Email":
		return customer.Email, nil
	case "AdressStreet":
		return customer.Address.Street, nil
	case "AdressCity":
		return customer.Address.City, nil
	case "AdressState":
		return customer.Address.State, nil
	case "AdressPostalCode":
		return customer.Address.PostalCode, nil
	case "AdressCountry":
		return customer.Address.Country, nil
	case "CreatedAt":
		return customer.CreatedAt, nil

	default:
		return nil, fmt.Errorf("field '%s' does not exist", field)
	}
}
