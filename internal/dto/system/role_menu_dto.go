package system

// AddRoleMenuDto 添加菜单角色关联请求参数
type AddRoleMenuDto struct {
	Id     int64 `json:"id"`     // 主键
	RoleId int64 `json:"roleId"` // 角色ID
	MenuId int64 `json:"menuId"` // 菜单ID
}

// DeleteRoleMenuDto 删除菜单角色关联请求参数
type DeleteRoleMenuDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateRoleMenuDto 修改菜单角色关联请求参数
type UpdateRoleMenuDto struct {
	Id     int64 `json:"id"`     // 主键
	RoleId int64 `json:"roleId"` // 角色ID
	MenuId int64 `json:"menuId"` // 菜单ID
}

// UpdateRoleMenuStatusDto 修改菜单角色关联状态请求参数
type UpdateRoleMenuStatusDto struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryRoleMenuDetailDto 查询菜单角色关联详情请求参数
type QueryRoleMenuDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryRoleMenuListDto 查询菜单角色关联列表请求参数
type QueryRoleMenuListDto struct {
	Id     int64 `json:"id"`     // 主键
	RoleId int64 `json:"roleId"` // 角色ID
	MenuId int64 `json:"menuId"` // 菜单ID
}
