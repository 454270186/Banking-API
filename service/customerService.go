package service

import "RESTful/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func NewCustomerService(repository domain.CustomerRepo) CustomerService {
	return DefaultCustomerService{repository}
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return d.repo.FindAll()
}
