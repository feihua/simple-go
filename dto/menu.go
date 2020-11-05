package dto

type MenuDto struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
}
