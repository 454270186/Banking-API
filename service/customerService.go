package service

import (
	"RESTful/domain"
	"RESTful/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func NewCustomerService(repository domain.CustomerRepo) CustomerService {
	return DefaultCustomerService{repository}
}

func (d DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	return d.repo.FindAll(status)
}

func (d DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return d.repo.FindById(id)
}
