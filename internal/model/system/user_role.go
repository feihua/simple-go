package system

// UserRole 角色用户关联
type UserRole struct {
	Id     int64 `gorm:"column:id" json:"id"`          // 主键
	UserId int64 `gorm:"column:user_id" json:"userId"` // 用户ID
	RoleId int64 `gorm:"column:role_id" json:"roleId"` // 角色ID
}

func (model UserRole) TableName() string {
	return "sys_user_role"
}
