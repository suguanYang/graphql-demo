package model

import "time"

type DeptEmp struct {
	EmpNo    string    `json:"emp_no"`
	DeptNo   string    `json:"dept_no"`
	FromDate time.Time `json:"from_date" gorm:"type:date"`
	ToDate   time.Time `json:"to_date" gorm:"type:date"`
}
