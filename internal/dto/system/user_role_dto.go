package system

// AddUserRoleDto 添加角色用户关联请求参数
type AddUserRoleDto struct {
	UserId int64 `json:"userId"` // 用户ID
	RoleId int64 `json:"roleId"` // 角色ID
}

// DeleteUserRoleDto 删除角色用户关联请求参数
type DeleteUserRoleDto struct {
	Ids []int64 `json:"ids"`
}
