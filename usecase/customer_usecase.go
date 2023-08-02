package usecase

import (
	"fmt"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(payload model.Customer) error
	FindAllCustomer(requestPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error)
	FindByIdCustomer(id string) (model.Customer, error)
	FindByPhoneNumber(phone string) (model.Customer, error)
	UpdateCustomer(payload model.Customer) error
	DeleteCustomer(id string) error
}

type cutomerUseCase struct {
	repo repository.CutomerRepository
}

func (c *cutomerUseCase) RegisterNewCustomer(payload model.Customer) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number and addres are required fields")
	}

	isExistedPhone, _ := c.repo.GetPhoneNumber(payload.PhoneNumber)

	if isExistedPhone.PhoneNumber == payload.PhoneNumber {
		return fmt.Errorf("customer with phone number %s already exists", payload.PhoneNumber)
	}

	err := c.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new customer: %v", err)
	}
	return nil
}

func (c *cutomerUseCase) FindAllCustomer(requestPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error) {
	return c.repo.Paging(requestPaging)
}

func (c *cutomerUseCase) FindByIdCustomer(id string) (model.Customer, error) {
	customer, err := c.repo.Get(id)
	if err != nil {
		return model.Customer{}, fmt.Errorf("tidak ada customer dengan id %s", id)
	}
	return customer, nil
}

func (c *cutomerUseCase) FindByPhoneNumber(phone string) (model.Customer, error) {
	return c.repo.GetPhoneNumber(phone)
}
func (c *cutomerUseCase) UpdateCustomer(payload model.Customer) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number and addres are required fields")
	}

	isExistedPhone, _ := c.repo.GetPhoneNumber(payload.PhoneNumber)
	if isExistedPhone.PhoneNumber == payload.PhoneNumber && isExistedPhone.Id != payload.Id {
		return fmt.Errorf("customer with phone number %s already exists", payload.PhoneNumber)
	}

	err := c.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update customer %s: %v", payload.Name, err)
	}
	return nil
}

func (c *cutomerUseCase) DeleteCustomer(id string) error {
	customer, err := c.FindByIdCustomer(id)
	if err != nil {
		return fmt.Errorf("data with id %s not found", id)
	}

	err = c.repo.Delete(customer.Id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %v", err)
	}
	return nil
}

func NewCustomerUseCase(repo repository.CutomerRepository) CustomerUseCase {
	return &cutomerUseCase{repo: repo}
}
