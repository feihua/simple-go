package dto

type RoleDto struct {
	ID       int64 `gorm:"primary_key"`
	Username string
	Password string
}
