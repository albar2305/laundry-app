package usecase

import (
	"fmt"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/repository"
)

type EmployeeUseCase interface {
	RegisterNewEmployee(payload model.Employee) error
	FindAllEmployee(requestPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error)
	FindByIdEmployee(id string) (model.Employee, error)
	FindByPhoneNumber(phone string) (model.Employee, error)
	UpdateEmployee(payload model.Employee) error
	DeleteEmployee(id string) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

func (e *employeeUseCase) RegisterNewEmployee(payload model.Employee) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number and addres are required fields")
	}

	isExistedPhone, _ := e.repo.GetByPhone(payload.PhoneNumber)

	if isExistedPhone.PhoneNumber == payload.PhoneNumber {
		return fmt.Errorf("employee with phone number %s already exists", payload.PhoneNumber)
	}

	err := e.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new employee: %v", err)
	}
	return nil
}
func (e *employeeUseCase) FindByPhoneNumber(phone string) (model.Employee, error) {
	return e.repo.GetByPhone(phone)
}

func (e *employeeUseCase) FindAllEmployee(requestPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error) {
	return e.repo.Paging(requestPaging)
}

func (e *employeeUseCase) FindByIdEmployee(id string) (model.Employee, error) {
	employee, err := e.repo.GetById(id)
	if err != nil {
		return model.Employee{}, fmt.Errorf("tidak ada employee dengan id %s", id)
	}
	return employee, nil
}

func (e *employeeUseCase) UpdateEmployee(payload model.Employee) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number and addres are required fields")
	}

	isExistedPhone, _ := e.repo.GetByPhone(payload.PhoneNumber)
	if isExistedPhone.PhoneNumber == payload.PhoneNumber && isExistedPhone.Id != payload.Id {
		return fmt.Errorf("employee with phone number %s already exists", payload.PhoneNumber)
	}

	err := e.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update employee %s: %v", payload.Name, err)
	}
	return nil
}

func (e *employeeUseCase) DeleteEmployee(id string) error {
	employee, err := e.FindByIdEmployee(id)
	if err != nil {
		return fmt.Errorf("data with id %s not found", id)
	}

	err = e.repo.Delete(employee.Id)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %v", err)
	}
	return nil
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
