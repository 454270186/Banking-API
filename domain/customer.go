package domain

type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	Birthday string
	Status string
}

type CustomerRepo interface {
	FindAll() ([]Customer, error)
}