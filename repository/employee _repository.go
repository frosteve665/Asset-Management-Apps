package repository

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/utils/constant"
	"database/sql"
	"fmt"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	// BaseRepositoryPaging[model.Employee]
}

type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) Create(payload model.Employee) error {
	_, err := e.db.Exec(constant.EMPLOYEE_INSERT, payload.Id, payload.Name, payload.Gender, payload.Address, payload.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

// func (e *employeeRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]model.Employee, dto.Paging, error) {
// 	var paginationQuery dto.PaginationQuery
// 	paginationQuery = common.GetPaginationParams(requestPaging)
// 	querySelect := "SELECT id, name, gender, phone_number, address FROM employee"
// 	if query[0] != "" {
// 		querySelect += ` WHERE name ilike '%` + query[0] + `%'`
// 	}
// 	querySelect += ` LIMIT $1 OFFSET $2`
// 	rows, err := e.db.Query(querySelect, paginationQuery.Take, paginationQuery.Skip)
// 	if err != nil {
// 		return nil, dto.Paging{}, err
// 	}
// 	var employees []model.Employee
// 	for rows.Next() {
// 		var employee model.Employee
// 		err := rows.Scan(&employee.Id, &employee.Gender, &employee.Name, &employee.PhoneNumber, &employee.Address)
// 		if err != nil {
// 			return nil, dto.Paging{}, err
// 		}
// 		employees = append(employees, employee)
// 	}

// 	// count total rows
// 	var totalRows int
// 	row := e.db.QueryRow("SELECT COUNT(*) FROM employee")
// 	err = row.Scan(&totalRows)
// 	if err != nil {
// 		return nil, dto.Paging{}, err
// 	}
// 	return employees, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
// }

func (e *employeeRepository) List() ([]model.Employee, error) {
	rows, err := e.db.Query(constant.EMPLOYEE_LIST)
	if err != nil {
		return nil, err
	}
	var employees []model.Employee

	for rows.Next() {
		var employee model.Employee
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Gender, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *employeeRepository) Get(id string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.QueryRow(constant.EMPLOYEE_GET, id).Scan(
		&employee.Id,
		&employee.Name,
		&employee.Gender,
		&employee.PhoneNumber,
		&employee.Address,
	)
	if err != nil {
		return model.Employee{}, fmt.Errorf("error get employee : %s ", err.Error())
	}
	return employee, nil
}

func (e *employeeRepository) Update(payload model.Employee) error {
	_, err := e.db.Exec(constant.EMPLOYEE_UPDATE, payload.Name, payload.Gender, payload.PhoneNumber, payload.Address, payload.Id)
	if err != nil {
		return fmt.Errorf("error update employee : %s ", err.Error())
	}
	return nil
}

func (e *employeeRepository) Delete(id string) error {
	_, err := e.db.Exec(constant.EMPLOYEE_DELETE, id)
	if err != nil {
		return fmt.Errorf("error delete employee : %s ", err.Error())
	}
	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}
