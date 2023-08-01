package repository

import (
	"database/sql"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/utils/common"
)

type CutomerRepository interface {
	BaseRepository[model.Customer]
	GetByPhone(phone string) (model.Customer, error)
	BaseRepositoryPaging[model.Customer]
}

type cutomerRepository struct {
	db *sql.DB
}

func (c *cutomerRepository) Create(payload model.Customer) error {
	_, err := c.db.Exec("INSERT INTO customer (id,name,phone_number,address) VALUES ($1,$2,$3,$4)", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}
func (c *cutomerRepository) List() ([]model.Customer, error) {
	rows, err := c.db.Query("SELECT id,name,phone_number,address FROM customer")
	if err != nil {
		return nil, err
	}

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
func (c *cutomerRepository) GetById(id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow("SELECT id,name,phone_number,address FROM customer WHERE id=$1", id).Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *cutomerRepository) Update(payload model.Customer) error {
	_, err := c.db.Exec("UPDATE customer SET name= $2, phone_number= $3, address=$4 WHERE id= $1", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}
func (c *cutomerRepository) GetByPhone(phone string) (model.Customer, error) {
	var cutomer model.Customer
	err := c.db.QueryRow("SELECT id,name,phone_number,address FROM customer WHERE phone_number = $1", phone).Scan(&cutomer.Id, &cutomer.Name, &cutomer.PhoneNumber, &cutomer.Address)
	if err != nil {
		return model.Customer{}, err
	}
	return cutomer, nil
}

func (c *cutomerRepository) Delete(id string) error {
	_, err := c.db.Exec("DELETE FROM customer WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (c *cutomerRepository) Paging(requestPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error) {
	paginationQuery := common.GetPaginationParams(requestPaging)
	rows, err := c.db.Query("SELECT id,name,phone_number,address FROM customer LIMIT $1 OFFSET $2", paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		customers = append(customers, customer)
	}

	var totalRows int
	row := c.db.QueryRow("SELECT COUNT(*) FROM customer")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	return customers, common.Pagination(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func NewCustomerRepository(db *sql.DB) CutomerRepository {
	return &cutomerRepository{
		db: db,
	}
}
