package model

type Department struct {
	DeptNo   string `json:"dept_no" gorm:"primary_key"`
	DeptName string `json:"dept_name" gorm:"unique_index"`
}
