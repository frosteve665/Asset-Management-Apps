package usecase

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/repository"
	"fmt"
)

type EmployeeUseCase interface {
	RegisterNewEmployee(payload model.Employee) error
	FindAllEmployeeList() ([]model.Employee, error)
	FindEmployeeById(id string) (model.Employee, error)
	UpdateEmployee(payload model.Employee) error
	DeleteEmployee(id string) error
	// FindAllEmployee(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.Employee, dto.Paging, error)
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

func (e *employeeUseCase) RegisterNewEmployee(payload model.Employee) error {
	//check attribute nama dan phoneNumber tidak boleh kosong
	if payload.Name == "" || payload.Gender == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, gender, Phone Number, Address is required")
	}
	err := e.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create Employee : %s", err.Error())
	}
	return nil
}

func (e *employeeUseCase) FindAllEmployeeList() ([]model.Employee, error) {
	return e.repo.List()
}

func (e *employeeUseCase) FindEmployeeById(id string) (model.Employee, error) {
	return e.repo.Get(id)
}

func (e *employeeUseCase) UpdateEmployee(payload model.Employee) error {
	err := e.repo.Update(payload)

	if err != nil {
		return err
	}
	return e.repo.Update(payload)
}

func (e *employeeUseCase) DeleteEmployee(id string) error {
	_, err := e.FindEmployeeById(id)
	if err != nil {
		return err
	}
	return e.repo.Delete(id)
}

// func (e *employeeUseCase) FindAllEmployee(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.Employee, dto.Paging, error) {
// 	return e.repo.Paging(requesPaging, byNameEmpl)
// }

func NewEmployeeUseCase(empRepo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{
		repo: empRepo,
	}
}
