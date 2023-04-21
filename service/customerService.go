package service

import (
	"RESTful/domain"
	"RESTful/dto"
	"RESTful/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func NewCustomerService(repository domain.CustomerRepo) CustomerService {
	return DefaultCustomerService{repository}
}

func (d DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, appErr := d.repo.FindAll(status)
	if appErr != nil {
		return nil, appErr
	}

	cusResponses := make([]dto.CustomerResponse, 0)
	for _, customer := range customers{
		response := customer.ToDTO()
		cusResponses = append(cusResponses, response)
	}

	return cusResponses, nil
}

func (d DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, appErr := d.repo.FindById(id)
	if appErr != nil {
		return nil, appErr
	}

	response := c.ToDTO()

	return &response, nil
}
