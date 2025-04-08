package system

import "time"

// AddUserDto 添加用户信息请求参数
type AddUserDto struct {
	Id            int64     `json:"id"`            // 主键
	Mobile        string    `json:"mobile"`        // 手机号码
	UserName      string    `json:"userName"`      // 用户账号
	NickName      string    `json:"nickName"`      // 用户昵称
	UserType      string    `json:"userType"`      // 用户类型（00系统用户）
	Avatar        string    `json:"avatar"`        // 头像路径
	Email         string    `json:"email"`         // 用户邮箱
	Password      string    `json:"password"`      // 密码
	Status        int32     `json:"status"`        // 状态(1:正常，0:禁用)
	DeptId        int64     `json:"deptId"`        // 部门ID
	LoginIp       string    `json:"loginIp"`       // 最后登录IP
	LoginDate     time.Time `json:"loginDate"`     // 最后登录时间
	LoginBrowser  string    `json:"loginBrowser"`  // 浏览器类型
	LoginOs       string    `json:"loginOs"`       // 操作系统
	PwdUpdateDate time.Time `json:"pwdUpdateDate"` // 密码最后更新时间
	Remark        string    `json:"remark"`        // 备注
	DelFlag       int32     `json:"delFlag"`       // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `json:"createBy"`      // 创建者
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateBy      string    `json:"updateBy"`      // 更新者
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}

// DeleteUserDto 删除用户信息请求参数
type DeleteUserDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateUserDto 修改用户信息请求参数
type UpdateUserDto struct {
	Id            int64     `json:"id"`            // 主键
	Mobile        string    `json:"mobile"`        // 手机号码
	UserName      string    `json:"userName"`      // 用户账号
	NickName      string    `json:"nickName"`      // 用户昵称
	UserType      string    `json:"userType"`      // 用户类型（00系统用户）
	Avatar        string    `json:"avatar"`        // 头像路径
	Email         string    `json:"email"`         // 用户邮箱
	Password      string    `json:"password"`      // 密码
	Status        int32     `json:"status"`        // 状态(1:正常，0:禁用)
	DeptId        int64     `json:"deptId"`        // 部门ID
	LoginIp       string    `json:"loginIp"`       // 最后登录IP
	LoginDate     time.Time `json:"loginDate"`     // 最后登录时间
	LoginBrowser  string    `json:"loginBrowser"`  // 浏览器类型
	LoginOs       string    `json:"loginOs"`       // 操作系统
	PwdUpdateDate time.Time `json:"pwdUpdateDate"` // 密码最后更新时间
	Remark        string    `json:"remark"`        // 备注
	DelFlag       int32     `json:"delFlag"`       // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `json:"createBy"`      // 创建者
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateBy      string    `json:"updateBy"`      // 更新者
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}

// UpdateUserStatusDto 修改用户信息状态请求参数
type UpdateUserStatusDto struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryUserDetailDto 查询用户信息详情请求参数
type QueryUserDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryUserListDto 查询用户信息列表请求参数
type QueryUserListDto struct {
	Id            int64     `json:"id"`            // 主键
	Mobile        string    `json:"mobile"`        // 手机号码
	UserName      string    `json:"userName"`      // 用户账号
	NickName      string    `json:"nickName"`      // 用户昵称
	UserType      string    `json:"userType"`      // 用户类型（00系统用户）
	Avatar        string    `json:"avatar"`        // 头像路径
	Email         string    `json:"email"`         // 用户邮箱
	Password      string    `json:"password"`      // 密码
	Status        int32     `json:"status"`        // 状态(1:正常，0:禁用)
	DeptId        int64     `json:"deptId"`        // 部门ID
	LoginIp       string    `json:"loginIp"`       // 最后登录IP
	LoginDate     time.Time `json:"loginDate"`     // 最后登录时间
	LoginBrowser  string    `json:"loginBrowser"`  // 浏览器类型
	LoginOs       string    `json:"loginOs"`       // 操作系统
	PwdUpdateDate time.Time `json:"pwdUpdateDate"` // 密码最后更新时间
	Remark        string    `json:"remark"`        // 备注
	DelFlag       int32     `json:"delFlag"`       // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `json:"createBy"`      // 创建者
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateBy      string    `json:"updateBy"`      // 更新者
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}
