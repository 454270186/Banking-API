package domain

import (
	"RESTful/dto"
	"RESTful/errs"
)

type Customer struct {
	Id       string
	Name     string
	City     string
	Zipcode  string
	Birthday string
	Status   string
}

type CustomerRepo interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}

func (c Customer) StatusAsText() string {
	statusAsText := "active"
	if c.Status == "false" {
		statusAsText = "inactive"
	}

	return statusAsText
}
 
func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id: c.Id,
		Name: c.Name,
		City: c.City,
		Zipcode: c.Zipcode,
		Birthday: c.Birthday,
		Status: c.StatusAsText(),
	}
}