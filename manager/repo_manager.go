package manager

import "github.com/albar2305/enigma-laundry-apps/repository"

type RepoManager interface {
	UomRepo() repository.UomRepository
	ProductRepo() repository.ProductRepository
	CustomerRepo() repository.CutomerRepository
	EmployeeRepo() repository.EmployeeRepository
	BillRepo() repository.BillRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

// UserRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

// BillRepo implements RepoManager.
func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.Conn())
}

// CustomerRepo implements RepoManager.
func (r *repoManager) CustomerRepo() repository.CutomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

// ProductRepo implements RepoManager.
func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.Conn())
}

// UomRepo implements RepoManager.
func (r *repoManager) UomRepo() repository.UomRepository {
	return repository.NewUomRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
