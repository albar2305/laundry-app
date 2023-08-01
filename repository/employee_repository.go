package repository

import (
	"database/sql"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/utils/common"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	GetByPhone(phone string) (model.Employee, error)
	BaseRepositoryPaging[model.Employee]
}

type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) Create(payload model.Employee) error {
	_, err := e.db.Exec("INSERT INTO employee (id,name,phone_number,address) VALUES ($1,$2,$3,$4)", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}
func (e *employeeRepository) List() ([]model.Employee, error) {
	rows, err := e.db.Query("SELECT id,name,phone_number,address FROM employee")
	if err != nil {
		return nil, err
	}

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
func (e *employeeRepository) GetById(id string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.QueryRow("SELECT id,name,phone_number,address FROM employee WHERE id=$1", id).Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) Update(payload model.Employee) error {
	_, err := e.db.Exec("UPDATE employee SET name= $2, phone_number= $3, address=$4 WHERE id= $1", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}
func (e *employeeRepository) GetByPhone(phone string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.QueryRow("SELECT id,name,phone_number,address FROM employee WHERE phone_number = $1", phone).Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) Paging(requestPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error) {
	paginationQuery := common.GetPaginationParams(requestPaging)
	rows, err := e.db.Query("SELECT id,name,phone_number,address FROM employee LIMIT $1 OFFSET $2", paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		employees = append(employees, employee)
	}

	var totalRows int
	row := e.db.QueryRow("SELECT COUNT(*) FROM employee")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	return employees, common.Pagination(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func (e *employeeRepository) Delete(id string) error {
	_, err := e.db.Exec("DELETE FROM employee WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}
