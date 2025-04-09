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

// QueryUserRoleListDto 查询角色用户关联列表请求参数
type QueryUserRoleListDto struct {
	PageNo   int   `json:"pageNo" default:"1"`    // 第几页
	PageSize int   `json:"pageSize" default:"10"` // 每页的数量
	Id       int64 `json:"id"`                    // 主键
	UserId   int64 `json:"userId"`                // 用户ID
	RoleId   int64 `json:"roleId"`                // 角色ID
}
