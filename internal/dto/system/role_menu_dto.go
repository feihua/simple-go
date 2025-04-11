package system

// AddRoleMenuDto 添加菜单角色关联请求参数
type AddRoleMenuDto struct {
	RoleId int64 `json:"roleId"` // 角色ID
	MenuId int64 `json:"menuId"` // 菜单ID
}

// DeleteRoleMenuDto 删除菜单角色关联请求参数
type DeleteRoleMenuDto struct {
	Ids []int64 `json:"ids"`
}

// QueryRoleMenuDetailDto 查询菜单角色关联详情请求参数
type QueryRoleMenuDetailDto struct {
	Id int64 `json:"id"` // id
}
