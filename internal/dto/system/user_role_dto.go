package system

// AddUserRoleDto 添加角色用户关联请求参数
type AddUserRoleDto struct {
	Id     int64 `json:"id"`     // 主键
	UserId int64 `json:"userId"` // 用户ID
	RoleId int64 `json:"roleId"` // 角色ID
}

// DeleteUserRoleDto 删除角色用户关联请求参数
type DeleteUserRoleDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateUserRoleDto 修改角色用户关联请求参数
type UpdateUserRoleDto struct {
	Id     int64 `json:"id"`     // 主键
	UserId int64 `json:"userId"` // 用户ID
	RoleId int64 `json:"roleId"` // 角色ID
}

// UpdateUserRoleStatusDto 修改角色用户关联状态请求参数
type UpdateUserRoleStatusDto struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryUserRoleDetailDto 查询角色用户关联详情请求参数
type QueryUserRoleDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryUserRoleListDto 查询角色用户关联列表请求参数
type QueryUserRoleListDto struct {
	Id     int64 `json:"id"`     // 主键
	UserId int64 `json:"userId"` // 用户ID
	RoleId int64 `json:"roleId"` // 角色ID
}
