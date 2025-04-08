package system

// RoleMenu 菜单角色关联
type RoleMenu struct {
	Id     int64 `gorm:"column:id" json:"id"`          // 主键
	RoleId int64 `gorm:"column:role_id" json:"roleId"` // 角色ID
	MenuId int64 `gorm:"column:menu_id" json:"menuId"` // 菜单ID
}

func (model RoleMenu) TableName() string {
	return "sys_role_menu"
}
