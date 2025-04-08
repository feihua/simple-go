package system

// AddUserRoleReqVo 添加角色用户关联请求参数
type AddUserRoleReqVo struct {
	Id     int64 `json:"id"`     // 主键
	UserId int64 `json:"userId"` // 用户ID
	RoleId int64 `json:"roleId"` // 角色ID
}

// DeleteUserRoleReqVo 删除角色用户关联请求参数
type DeleteUserRoleReqVo struct {
	Ids []int64 `json:"ids"`
}

// UpdateUserRoleReqVo 修改角色用户关联请求参数
type UpdateUserRoleReqVo struct {
	Id     int64 `json:"id"`     // 主键
	UserId int64 `json:"userId"` // 用户ID
	RoleId int64 `json:"roleId"` // 角色ID
}

// UpdateUserRoleStatusReqVo 修改角色用户关联状态请求参数
type UpdateUserRoleStatusReqVo struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryUserRoleDetailReqVo 查询角色用户关联详情请求参数
type QueryUserRoleDetailReqVo struct {
	Id int64 `json:"id"` // id
}

// QueryUserRoleListReqVo 查询角色用户关联列表请求参数
type QueryUserRoleListReqVo struct {
	Id     int64 `json:"id"`     // 主键
	UserId int64 `json:"userId"` // 用户ID
	RoleId int64 `json:"roleId"` // 角色ID
}
