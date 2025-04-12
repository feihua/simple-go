package system

// AddRoleReqVo 添加角色信息请求参数
type AddRoleReqVo struct {
	RoleName  string `json:"roleName" binding:"required"`  // 名称
	RoleKey   string `json:"roleKey" binding:"required"`   // 角色权限字符串
	DataScope int32  `json:"dataScope" binding:"required"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status    int32  `json:"status" binding:"required"`    // 状态(1:正常，0:禁用)
	Remark    string `json:"remark" binding:"required"`    // 备注

}

// DeleteRoleReqVo 删除角色信息请求参数
type DeleteRoleReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdateRoleReqVo 修改角色信息请求参数
type UpdateRoleReqVo struct {
	Id        int64  `json:"id" binding:"required"`        // 主键
	RoleName  string `json:"roleName" binding:"required"`  // 名称
	RoleKey   string `json:"roleKey" binding:"required"`   // 角色权限字符串
	DataScope int32  `json:"dataScope" binding:"required"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status    int32  `json:"status" binding:"required"`    // 状态(1:正常，0:禁用)
	Remark    string `json:"remark" binding:"required"`    // 备注

}

// UpdateRoleStatusReqVo 修改角色信息状态请求参数
type UpdateRoleStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryRoleDetailReqVo 查询角色信息详情请求参数
type QueryRoleDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryRoleListReqVo 查询角色信息列表请求参数
type QueryRoleListReqVo struct {
	PageNo    int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize  int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	RoleName  string `json:"roleName" `                                // 名称
	RoleKey   string `json:"roleKey" `                                 // 角色权限字符串
	DataScope *int32 `json:"dataScope" `                               // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status    *int32 `json:"status" `                                  // 状态(1:正常，0:禁用)

}

type QueryRoleMenuListReq struct {
	RoleId int64 `json:"roleId" binding:"required"` // 角色Id
}

type UpdateRoleMenuReq struct {
	RoleId  int64   `json:"roleId" binding:"required"`  // 角色Id
	MenuIds []int64 `json:"menuIds" binding:"required"` // 菜单Id
}

// QueryRoleUserListReqVo 查询角色用户信息列表请求参数
type QueryRoleUserListReqVo struct {
	PageNo   int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	RoleId   int64  `json:"RoleId" binding:"required"`                // 主键
	Mobile   string `json:"mobile" `                                  // 手机号码
	UserName string `json:"userName" "`                               // 用户账号
}

type AuthUserReq struct {
	RoleId int64 `json:"roleId" binding:"required"` // 角色Id
	UserId int64 `json:"userId" binding:"required"` // 用户Id
}

type BatchAuthUserReq struct {
	RoleId  int64   `json:"roleId" binding:"required"`  // 角色Id
	UserIds []int64 `json:"userIds" binding:"required"` // 用户Id
}
