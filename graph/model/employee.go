package model

import "time"

type Employee struct {
	EmpNo     string    `json:"emp_no" gorm:"primary_key"`
	BirthDate time.Time `json:"birth_date" gorm:"type:date"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gener" gorm:"type:ENUM('M', 'F')"`
	HireDate  time.Time `json:"hire_date" gorm:"type:date"`
}
