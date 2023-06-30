package services

import (
	"github.com/KhoirulAziz99/mnc/internal/models"
	"github.com/KhoirulAziz99/mnc/internal/repository"
)
//saya jelaskan secara garia besar, package ini adalah services, biasanya untuk bisnis logic, berhubung tidak adabisnis logic jadi tiap methodnya hanya mengembalikan repository
type CustomerServices interface {
	Register(newCustomer models.Customer) error
	GetAll()([]models.Customer, error)
	GetByEmail(email string) (*models.Customer, error)

}

type customerServices struct {
	customerRepository repository.CustomerRepository
}
func NewCustomerServices(repo repository.CustomerRepository) *customerServices {
	return &customerServices{customerRepository: repo}
}

func (c *customerServices) Register(newCustomer models.Customer) error {
	return c.customerRepository.Created(newCustomer)
}

func (c *customerServices) GetAll()([]models.Customer, error) {
	return c.customerRepository.FindAll()
}

func (c *customerServices) GetByEmail(email string) (*models.Customer, error) {
	return c.customerRepository.FindByEmail(email)
}