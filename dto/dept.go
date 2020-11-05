package dto

type DeptDto struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
}
