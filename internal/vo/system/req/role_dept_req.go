package system

// AddRoleDeptReqVo 添加角色和部门关联请求参数
type AddRoleDeptReqVo struct {
	RoleId int64 `json:"roleId"` // 角色id
	DeptId int64 `json:"deptId"` // 部门id
}

// DeleteRoleDeptReqVo 删除角色和部门关联请求参数
type DeleteRoleDeptReqVo struct {
	Ids []int64 `json:"ids"`
}

// UpdateRoleDeptReqVo 修改角色和部门关联请求参数
type UpdateRoleDeptReqVo struct {
	RoleId int64 `json:"roleId"` // 角色id
	DeptId int64 `json:"deptId"` // 部门id
}

// UpdateRoleDeptStatusReqVo 修改角色和部门关联状态请求参数
type UpdateRoleDeptStatusReqVo struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryRoleDeptDetailReqVo 查询角色和部门关联详情请求参数
type QueryRoleDeptDetailReqVo struct {
	Id int64 `json:"id"` // id
}

// QueryRoleDeptListReqVo 查询角色和部门关联列表请求参数
type QueryRoleDeptListReqVo struct {
	RoleId int64 `json:"roleId"` // 角色id
	DeptId int64 `json:"deptId"` // 部门id
}
