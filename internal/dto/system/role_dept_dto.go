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

// QueryRoleDeptDetailDto 查询角色和部门关联详情请求参数
type QueryRoleDeptDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryRoleDeptListDto 查询角色和部门关联列表请求参数
type QueryRoleDeptListDto struct {
	PageNo   int   `json:"pageNo" default:"1"`    // 第几页
	PageSize int   `json:"pageSize" default:"10"` // 每页的数量
	RoleId   int64 `json:"roleId"`                // 角色id
	DeptId   int64 `json:"deptId"`                // 部门id
}
