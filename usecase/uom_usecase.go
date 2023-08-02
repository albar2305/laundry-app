package usecase

import (
	"fmt"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/repository"
)

type UomUseCase interface {
	RegisterNewUom(payload model.Uom) error
	FindAllUom() ([]model.Uom, error)
	FindByIdUom(id string) (model.Uom, error)
	UpdateUom(payload model.Uom) error
	DeleteUom(id string) error
}

type uomUseCase struct {
	repo repository.UomRepository
}

func (u *uomUseCase) RegisterNewUom(payload model.Uom) error {
	if payload.Name == "" {
		return fmt.Errorf("name required fields")
	}

	isExistedUom, _ := u.repo.GetByName(payload.Name)
	if isExistedUom.Name == payload.Name {
		return fmt.Errorf("uom with name %s already exists", payload.Name)
	}

	err := u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new uom: %v", err)
	}
	return nil
}

func (u *uomUseCase) FindAllUom() ([]model.Uom, error) {
	return u.repo.List()
}

func (u *uomUseCase) FindByIdUom(id string) (model.Uom, error) {
	return u.repo.Get(id)
}

func (u *uomUseCase) UpdateUom(payload model.Uom) error {
	if payload.Name == "" {
		return fmt.Errorf("name is required field")
	}

	isExistUom, _ := u.repo.GetByName(payload.Name)
	if isExistUom.Name == payload.Name && isExistUom.Id != payload.Id {
		return fmt.Errorf("uom with name %s already exists", payload.Name)
	}

	err := u.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom %s: %v", payload.Name, err)
	}
	return nil
}

func (u *uomUseCase) DeleteUom(id string) error {
	uom, err := u.FindByIdUom(id)
	if err != nil {
		return fmt.Errorf("data with id %s not found", id)
	}

	err = u.repo.Delete(uom.Id)
	if err != nil {
		return fmt.Errorf("failed to delete uom: %v", err)
	}
	return nil
}

func NewUomUseCase(repo repository.UomRepository) UomUseCase {
	return &uomUseCase{repo: repo}
}
