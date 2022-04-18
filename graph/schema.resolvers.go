package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/suguanyang/graphql-go/graph/generated"
	"github.com/suguanyang/graphql-go/graph/model"
)

func (r *mutationResolver) AddDepartment(ctx context.Context, deptName string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddEmployee(ctx context.Context, birthDate time.Time, firstName string, lastName string, gender string, hireDate time.Time) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Departments(ctx context.Context, limit *int, offset *int) ([]*model.Department, error) {
	return r.Repository.GetDepartments(limit, offset), nil
}

func (r *queryResolver) DepartmentByDeptNo(ctx context.Context, deptNo string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DepartmentByEmpNo(ctx context.Context, empNo string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employees(ctx context.Context, limit *int, offset *int) ([]*model.Employee, error) {
	return r.Repository.GetEmployees(limit, offset), nil
}

func (r *queryResolver) EmployeeByDeptNo(ctx context.Context, deptNo string) ([]*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EmployeeByEmpNo(ctx context.Context, empNo string) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Salaries(ctx context.Context, limit *int, offset *int) ([]*model.Salary, error) {
	return r.Repository.GetSalaries(limit, offset), nil
}

func (r *queryResolver) SalaryByEmpNo(ctx context.Context, empNo string) (*model.Salary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DeptManagers(ctx context.Context, limit *int, offset *int) ([]*model.DeptManager, error) {
	return r.Repository.GetDepartmentMangers(limit, offset), nil
}

func (r *queryResolver) DeptManagerByEmpNo(ctx context.Context, empNo string) ([]*model.DeptManager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DeptManagerByDeptNo(ctx context.Context, deptNo string) ([]*model.DeptManager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Titles(ctx context.Context, limit *int, offset *int) ([]*model.Title, error) {
	return r.Repository.GetTitles(limit, offset), nil
}

func (r *queryResolver) TitleByEmpNo(ctx context.Context, empNo string) (*model.Title, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *salaryResolver) FromDate(ctx context.Context, obj *model.Salary) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *salaryResolver) ToDate(ctx context.Context, obj *model.Salary) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *titleResolver) FromDate(ctx context.Context, obj *model.Title) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *titleResolver) ToDate(ctx context.Context, obj *model.Title) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Salary returns generated.SalaryResolver implementation.
func (r *Resolver) Salary() generated.SalaryResolver { return &salaryResolver{r} }

// Title returns generated.TitleResolver implementation.
func (r *Resolver) Title() generated.TitleResolver { return &titleResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type salaryResolver struct{ *Resolver }
type titleResolver struct{ *Resolver }
