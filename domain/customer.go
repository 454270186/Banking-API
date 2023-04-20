package domain

import "RESTful/errs"

type Customer struct {
	Id       string
	Name     string
	City     string
	Zipcode  string
	Birthday string
	Status   string
}

type CustomerRepo interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}