package model

type Salary struct {
	EmpNo    string `json:"emp_no" gorm:"primary_key"`
	Salary   string `json:"salary"`
	FromDate string `json:"from_date" gorm:"type:date"`
	ToDate   string `json:"to_date" gorm:"type:date"`
}
