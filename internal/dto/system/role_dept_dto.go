package system

// AddRoleDeptDto 添加角色和部门关联请求参数
type AddRoleDeptDto struct {
	RoleId int64 `json:"roleId"` // 角色id
	DeptId int64 `json:"deptId"` // 部门id
}

// DeleteRoleDeptDto 删除角色和部门关联请求参数
type DeleteRoleDeptDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateRoleDeptDto 修改角色和部门关联请求参数
type UpdateRoleDeptDto struct {
	RoleId int64 `json:"roleId"` // 角色id
	DeptId int64 `json:"deptId"` // 部门id
}

// UpdateRoleDeptStatusDto 修改角色和部门关联状态请求参数
type UpdateRoleDeptStatusDto struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryRoleDeptDetailDto 查询角色和部门关联详情请求参数
type QueryRoleDeptDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryRoleDeptListDto 查询角色和部门关联列表请求参数
type QueryRoleDeptListDto struct {
	RoleId int64 `json:"roleId"` // 角色id
	DeptId int64 `json:"deptId"` // 部门id
}
