package repository

import (
	"github.com/suguanyang/graphql-go/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetEmployees(limit *int, offset *int) []*model.Employee
	GetDepartments(limit *int, offset *int) []*model.Department
	GetSalaries(limit *int, offset *int) []*model.Salary
	GetTitles(limit *int, offset *int) []*model.Title
	GetDepartmentMangers(limit *int, offset *int) []*model.DeptManager
	GetEmployeeByEmpNo(empNo string) *model.Employee
}

type EmployeeDatabase struct {
	db *gorm.DB
}

func connect() *gorm.DB {
	dsn := "root:1234@tcp(127.0.0.1:3306)/employees?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewEmployeeRepository() *EmployeeDatabase {
	return &EmployeeDatabase{
		db: connect(),
	}
}

func (source *EmployeeDatabase) GetEmployees(limit *int, offset *int) []*model.Employee {
	employees := make([]*model.Employee, 0)
	pagination := NewPagination(limit, offset)
	if err := source.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&employees).Error; err != nil {
		panic(err)
	}
	return employees
}

func (source *EmployeeDatabase) GetEmployeeByEmpNo(empNo string) (employee *model.Employee) {
	employee = &model.Employee{}
	if err := source.db.Where("emp_no = ?", empNo).First(employee).Error; err != nil {
		panic(err)
	}
	return
}

func (source *EmployeeDatabase) GetDepartments(limit *int, offset *int) []*model.Department {
	departments := make([]*model.Department, 0)
	pagination := NewPagination(limit, offset)
	if err := source.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(departments).Error; err != nil {
		panic(err)
	}
	return departments
}

func (source *EmployeeDatabase) GetSalaries(limit *int, offset *int) []*model.Salary {
	salaries := make([]*model.Salary, 0)
	pagination := NewPagination(limit, offset)
	if err := source.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(salaries).Error; err != nil {
		panic(err)
	}
	return salaries
}

func (source *EmployeeDatabase) GetTitles(limit *int, offset *int) []*model.Title {
	titles := make([]*model.Title, 0)
	pagination := NewPagination(limit, offset)
	if err := source.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(titles).Error; err != nil {
		panic(err)
	}
	return titles
}

func (source *EmployeeDatabase) GetDepartmentMangers(limit *int, offset *int) []*model.DeptManager {
	departmentMangers := make([]*model.DeptManager, 0)
	pagination := NewPagination(limit, offset)
	if err := source.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(departmentMangers).Error; err != nil {
		panic(err)
	}
	return departmentMangers
}
