package repository

import (
	"database/sql"

	"github.com/albar2305/enigma-laundry-apps/model"
)

type UomRepository interface {
	BaseRepository[model.Uom]
	GetByName(name string) (model.Uom, error)
}

type uomRepository struct {
	db *sql.DB
}

func (u *uomRepository) Create(payload model.Uom) error {
	_, err := u.db.Exec("INSERT INTO uom (id,name) VALUES ($1,$2)", payload.Id, payload.Name)
	if err != nil {
		return err
	}
	return nil
}
func (u *uomRepository) List() ([]model.Uom, error) {
	rows, err := u.db.Query("SELECT id,name FROM uom")
	if err != nil {
		return nil, err
	}

	var uoms []model.Uom
	for rows.Next() {
		var uom model.Uom
		err := rows.Scan(&uom.Id, &uom.Name)
		if err != nil {
			return nil, err
		}
		uoms = append(uoms, uom)
	}
	return uoms, nil
}
func (u *uomRepository) GetById(id string) (model.Uom, error) {
	var uom model.Uom
	err := u.db.QueryRow("SELECT id,name FROM uom WHERE id=$1", id).Scan(&uom.Id, &uom.Name)
	if err != nil {
		return model.Uom{}, err
	}
	return uom, nil
}

func (u *uomRepository) GetByName(name string) (model.Uom, error) {
	var uom model.Uom
	err := u.db.QueryRow("SELECT id,name FROM uom WHERE name ILIKE $1", "%"+name+"%").Scan(&uom.Id, &uom.Name)
	if err != nil {
		return model.Uom{}, err
	}
	return uom, nil
}
func (u *uomRepository) Update(payload model.Uom) error {
	_, err := u.db.Exec("UPDATE uom SET name= $1 WHERE id= $2", payload.Name, payload.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *uomRepository) Delete(id string) error {
	_, err := u.db.Exec("DELETE FROM uom WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{
		db: db,
	}
}
