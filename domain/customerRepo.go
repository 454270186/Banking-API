package domain

type CustomerRepoStub struct {
	customers []Customer
}

func NewCustomerRepoStub() CustomerRepoStub {
	customers := []Customer{
		{"1001", "Xiaofei", "Chengdu", "666", "0418", "1"},
		{"1002", "xiaoxiao", "Chengdu", "666", "0419", "1"},
	}

	return CustomerRepoStub{customers}
}

// FindAll() returns the slice of customers
func (c CustomerRepoStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}



