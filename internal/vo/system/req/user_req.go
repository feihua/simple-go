package system

import "time"

// AddUserReqVo 添加用户信息请求参数
type AddUserReqVo struct {
	Id            int64     `json:"id" binding:"required"`            // 主键
	Mobile        string    `json:"mobile" binding:"required"`        // 手机号码
	UserName      string    `json:"userName" binding:"required"`      // 用户账号
	NickName      string    `json:"nickName" binding:"required"`      // 用户昵称
	UserType      string    `json:"userType" binding:"required"`      // 用户类型（00系统用户）
	Avatar        string    `json:"avatar" binding:"required"`        // 头像路径
	Email         string    `json:"email" binding:"required"`         // 用户邮箱
	Password      string    `json:"password" binding:"required"`      // 密码
	Status        int32     `json:"status" binding:"required"`        // 状态(1:正常，0:禁用)
	DeptId        int64     `json:"deptId" binding:"required"`        // 部门ID
	LoginIp       string    `json:"loginIp" binding:"required"`       // 最后登录IP
	LoginDate     time.Time `json:"loginDate" binding:"required"`     // 最后登录时间
	LoginBrowser  string    `json:"loginBrowser" binding:"required"`  // 浏览器类型
	LoginOs       string    `json:"loginOs" binding:"required"`       // 操作系统
	PwdUpdateDate time.Time `json:"pwdUpdateDate" binding:"required"` // 密码最后更新时间
	Remark        string    `json:"remark" binding:"required"`        // 备注
	DelFlag       int32     `json:"delFlag" binding:"required"`       // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `json:"createBy" binding:"required"`      // 创建者
	CreateTime    time.Time `json:"createTime" binding:"required"`    // 创建时间
	UpdateBy      string    `json:"updateBy" binding:"required"`      // 更新者
	UpdateTime    time.Time `json:"updateTime" binding:"required"`    // 更新时间
}

// DeleteUserReqVo 删除用户信息请求参数
type DeleteUserReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// UpdateUserReqVo 修改用户信息请求参数
type UpdateUserReqVo struct {
	Id            int64     `json:"id" binding:"required"`            // 主键
	Mobile        string    `json:"mobile" binding:"required"`        // 手机号码
	UserName      string    `json:"userName" binding:"required"`      // 用户账号
	NickName      string    `json:"nickName" binding:"required"`      // 用户昵称
	UserType      string    `json:"userType" binding:"required"`      // 用户类型（00系统用户）
	Avatar        string    `json:"avatar" binding:"required"`        // 头像路径
	Email         string    `json:"email" binding:"required"`         // 用户邮箱
	Password      string    `json:"password" binding:"required"`      // 密码
	Status        int32     `json:"status" binding:"required"`        // 状态(1:正常，0:禁用)
	DeptId        int64     `json:"deptId" binding:"required"`        // 部门ID
	LoginIp       string    `json:"loginIp" binding:"required"`       // 最后登录IP
	LoginDate     time.Time `json:"loginDate" binding:"required"`     // 最后登录时间
	LoginBrowser  string    `json:"loginBrowser" binding:"required"`  // 浏览器类型
	LoginOs       string    `json:"loginOs" binding:"required"`       // 操作系统
	PwdUpdateDate time.Time `json:"pwdUpdateDate" binding:"required"` // 密码最后更新时间
	Remark        string    `json:"remark" binding:"required"`        // 备注
	DelFlag       int32     `json:"delFlag" binding:"required"`       // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `json:"createBy" binding:"required"`      // 创建者
	CreateTime    time.Time `json:"createTime" binding:"required"`    // 创建时间
	UpdateBy      string    `json:"updateBy" binding:"required"`      // 更新者
	UpdateTime    time.Time `json:"updateTime" binding:"required"`    // 更新时间
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
	PageNo        int       `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize      int       `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	Id            int64     `json:"id" binding:"required"`                    // 主键
	Mobile        string    `json:"mobile" binding:"required"`                // 手机号码
	UserName      string    `json:"userName" binding:"required"`              // 用户账号
	NickName      string    `json:"nickName" binding:"required"`              // 用户昵称
	UserType      string    `json:"userType" binding:"required"`              // 用户类型（00系统用户）
	Avatar        string    `json:"avatar" binding:"required"`                // 头像路径
	Email         string    `json:"email" binding:"required"`                 // 用户邮箱
	Password      string    `json:"password" binding:"required"`              // 密码
	Status        int32     `json:"status" binding:"required"`                // 状态(1:正常，0:禁用)
	DeptId        int64     `json:"deptId" binding:"required"`                // 部门ID
	LoginIp       string    `json:"loginIp" binding:"required"`               // 最后登录IP
	LoginDate     time.Time `json:"loginDate" binding:"required"`             // 最后登录时间
	LoginBrowser  string    `json:"loginBrowser" binding:"required"`          // 浏览器类型
	LoginOs       string    `json:"loginOs" binding:"required"`               // 操作系统
	PwdUpdateDate time.Time `json:"pwdUpdateDate" binding:"required"`         // 密码最后更新时间
	Remark        string    `json:"remark" binding:"required"`                // 备注
	DelFlag       int32     `json:"delFlag" binding:"required"`               // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `json:"createBy" binding:"required"`              // 创建者
	CreateTime    time.Time `json:"createTime" binding:"required"`            // 创建时间
	UpdateBy      string    `json:"updateBy" binding:"required"`              // 更新者
	UpdateTime    time.Time `json:"updateTime" binding:"required"`            // 更新时间
}

// LoginReqVo 登录参数
type LoginReqVo struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserRoleRequest 更新用户角色参数
type UpdateUserRoleRequest struct {
	UserId int64   `json:"userId" binding:"required"`
	RoleId []int64 `json:"roleId" binding:"required"`
}
