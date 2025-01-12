package dao

type AddressDAO struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type CustomerDAO struct {
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Address AddressDAO `json:"address"`
}
