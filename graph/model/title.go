package model

type Title struct {
	EmpNo    string `json:"emp_no" gorm:"primary_key"`
	Title    string `json:"title"`
	FromDate string `json:"from_date" gorm:"type:date"`
	ToDate   string `json:"to_date" gorm:"type:date"`
}
