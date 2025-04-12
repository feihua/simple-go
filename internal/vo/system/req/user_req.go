package system

// AddUserReqVo 添加用户信息请求参数
type AddUserReqVo struct {
	Mobile   string  `json:"mobile" binding:"required"`   // 手机号码
	UserName string  `json:"userName" binding:"required"` // 用户账号
	NickName string  `json:"nickName" binding:"required"` // 用户昵称
	UserType string  `json:"userType" binding:"required"` // 用户类型（00系统用户）
	Avatar   string  `json:"avatar" binding:"required"`   // 头像路径
	Email    string  `json:"email" binding:"required"`    // 用户邮箱
	Password string  `json:"password" binding:"required"` // 密码
	Status   int32   `json:"status" binding:"required"`   // 状态(1:正常，0:禁用)
	DeptId   int64   `json:"deptId" binding:"required"`   // 部门ID
	Remark   string  `json:"remark" binding:"required"`   // 备注
	PostIds  []int64 `json:"postIds"`                     // 部门id
}

// DeleteUserReqVo 删除用户信息请求参数
type DeleteUserReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdateUserReqVo 修改用户信息请求参数
type UpdateUserReqVo struct {
	Id       int64   `json:"id" binding:"required"`       // 主键
	Mobile   string  `json:"mobile" binding:"required"`   // 手机号码
	UserName string  `json:"userName" binding:"required"` // 用户账号
	NickName string  `json:"nickName" binding:"required"` // 用户昵称
	UserType string  `json:"userType" binding:"required"` // 用户类型（00系统用户）
	Avatar   string  `json:"avatar" binding:"required"`   // 头像路径
	Email    string  `json:"email" binding:"required"`    // 用户邮箱
	Status   int32   `json:"status" binding:"required"`   // 状态(1:正常，0:禁用)
	DeptId   int64   `json:"deptId" binding:"required"`   // 部门ID
	Remark   string  `json:"remark" binding:"required"`   // 备注
	PostIds  []int64 `json:"postIds"`                     // 部门id
}

// UpdateUserStatusReqVo 修改用户信息状态请求参数
type UpdateUserStatusReqVo struct {
	Ids    []int64 `json:"ids" binding:"required"`    // id
	Status int32   `json:"status" binding:"required"` // 状态（0:关闭,1:正常 ）
}

// QueryUserDetailReqVo 查询用户信息详情请求参数
type QueryUserDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryUserListReqVo 查询用户信息列表请求参数
type QueryUserListReqVo struct {
	PageNo   int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	Mobile   string `json:"mobile" `                                  // 手机号码
	UserName string `json:"userName" `                                // 用户账号
	NickName string `json:"nickName" `                                // 用户昵称
	Status   *int32 `json:"status" `                                  // 状态(1:正常，0:禁用)
	DeptId   *int64 `json:"deptId" `                                  // 部门ID

}

// LoginReqVo 登录参数
type LoginReqVo struct {
	Account  string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// QueryUserRoleListReqVo 查询用户已分配角色请求参数
type QueryUserRoleListReqVo struct {
	PageNo   int   `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize int   `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	UserId   int64 `json:"userId" binding:"required"`
}

// UpdateUserRoleReqVo 更新用户角色参数
type UpdateUserRoleReqVo struct {
	UserId  int64   `json:"userId" binding:"required"`
	RoleIds []int64 `json:"roleIds" binding:"required"`
}
