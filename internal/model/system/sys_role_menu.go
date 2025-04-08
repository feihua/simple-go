package system

// SysRoleMenu 菜单角色关联表
type SysRoleMenu struct {
	Id     int64 `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"` // 主键
	RoleId int64 `gorm:"column:role_id;NOT NULL" json:"role_id"`         // 角色ID
	MenuId int64 `gorm:"column:menu_id;NOT NULL" json:"menu_id"`         // 菜单ID
}

func (m *SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
