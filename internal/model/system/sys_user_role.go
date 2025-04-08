package system

// SysUserRole 角色用户关联表
type SysUserRole struct {
	Id     int64 `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`   // 主键
	UserId int64 `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"` // 用户ID
	RoleId int64 `gorm:"column:role_id;NOT NULL" json:"role_id"`           // 角色ID
}

func (m *SysUserRole) TableName() string {
	return "sys_user_role"
}
