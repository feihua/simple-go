package system

// RoleMenu 菜单角色关联
type RoleMenu struct {
	Id     int64 `json:"id"`     // 主键
	RoleId int64 `json:"roleId"` // 角色ID
	MenuId int64 `json:"menuId"` // 菜单ID
}
