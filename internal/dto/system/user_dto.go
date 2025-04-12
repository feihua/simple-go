package system

import "time"

// AddUserDto 添加用户信息请求参数
type AddUserDto struct {
	Id       int64   `json:"id"`       // 主键
	Mobile   string  `json:"mobile"`   // 手机号码
	UserName string  `json:"userName"` // 用户账号
	NickName string  `json:"nickName"` // 用户昵称
	UserType string  `json:"userType"` // 用户类型（00系统用户）
	Avatar   string  `json:"avatar"`   // 头像路径
	Email    string  `json:"email"`    // 用户邮箱
	Password string  `json:"password"` // 密码
	Status   int32   `json:"status"`   // 状态(1:正常，0:禁用)
	DeptId   int64   `json:"deptId"`   // 部门ID
	Remark   string  `json:"remark"`   // 备注
	CreateBy string  `json:"createBy"` // 创建者
	PostIds  []int64 `json:"postIds"`  // 部门id
}

// DeleteUserDto 删除用户信息请求参数
type DeleteUserDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateUserDto 修改用户信息请求参数
type UpdateUserDto struct {
	Id            int64      `json:"id"`            // 主键
	Mobile        string     `json:"mobile"`        // 手机号码
	UserName      string     `json:"userName"`      // 用户账号
	NickName      string     `json:"nickName"`      // 用户昵称
	UserType      string     `json:"userType"`      // 用户类型（00系统用户）
	Avatar        string     `json:"avatar"`        // 头像路径
	Email         string     `json:"email"`         // 用户邮箱
	Password      string     `json:"password"`      // 密码
	Status        int32      `json:"status"`        // 状态(1:正常，0:禁用)
	DeptId        int64      `json:"deptId"`        // 部门ID
	LoginIp       string     `json:"loginIp"`       // 最后登录IP
	LoginDate     *time.Time `json:"loginDate"`     // 最后登录时间
	LoginBrowser  string     `json:"loginBrowser"`  // 浏览器类型
	LoginOs       string     `json:"loginOs"`       // 操作系统
	PwdUpdateDate *time.Time `json:"pwdUpdateDate"` // 密码最后更新时间
	Remark        string     `json:"remark"`        // 备注
	DelFlag       int32      `json:"delFlag"`       // 删除标志（0代表删除 1代表存在）
	CreateBy      string     `json:"createBy"`      // 创建者
	CreateTime    time.Time  `json:"createTime"`    // 创建时间
	UpdateBy      string     `json:"updateBy"`      // 更新者
	UpdateTime    time.Time  `json:"updateTime"`    // 更新时间
	PostIds       []int64    `json:"postIds"`       // 部门id
}

// UpdateUserStatusDto 修改用户信息状态请求参数
type UpdateUserStatusDto struct {
	Ids      []int64 `json:"ids"`      // id
	Status   int32   `json:"status"`   // 状态（0:关闭,1:正常 ）
	UpdateBy string  `json:"updateBy"` // 更新者
}

// QueryUserDetailDto 查询用户信息详情请求参数
type QueryUserDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryUserListDto 查询用户信息列表请求参数
type QueryUserListDto struct {
	PageNo   int    `json:"pageNo" default:"1"`    // 第几页
	PageSize int    `json:"pageSize" default:"10"` // 每页的数量
	Mobile   string `json:"mobile"`                // 手机号码
	UserName string `json:"userName"`              // 用户账号
	NickName string `json:"nickName"`              // 用户昵称
	Status   *int32 `json:"status"`                // 状态(1:正常，0:禁用)
	DeptId   *int64 `json:"deptId"`                // 部门ID
}

// QueryUserListDtoResp 查询用户信息列表响应参数
type QueryUserListDtoResp struct {
	Id            int64   `json:"id"`            // 主键
	Mobile        string  `json:"mobile"`        // 手机号码
	UserName      string  `json:"userName"`      // 用户账号
	NickName      string  `json:"nickName"`      // 用户昵称
	UserType      string  `json:"userType"`      // 用户类型（00系统用户）
	Avatar        string  `json:"avatar"`        // 头像路径
	Email         string  `json:"email"`         // 用户邮箱
	Password      string  `json:"password"`      // 密码
	Status        int32   `json:"status"`        // 状态(1:正常，0:禁用)
	DeptId        int64   `json:"deptId"`        // 部门ID
	LoginIp       string  `json:"loginIp"`       // 最后登录IP
	LoginDate     string  `json:"loginDate"`     // 最后登录时间
	LoginBrowser  string  `json:"loginBrowser"`  // 浏览器类型
	LoginOs       string  `json:"loginOs"`       // 操作系统
	PwdUpdateDate string  `json:"pwdUpdateDate"` // 密码最后更新时间
	Remark        string  `json:"remark"`        // 备注
	DelFlag       int32   `json:"delFlag"`       // 删除标志（0代表删除 1代表存在）
	CreateBy      string  `json:"createBy"`      // 创建者
	CreateTime    string  `json:"createTime"`    // 创建时间
	UpdateBy      string  `json:"updateBy"`      // 更新者
	UpdateTime    string  `json:"updateTime"`    // 更新时间
	PostIds       []int64 `json:"postIds"`       // 部门id
}

// LoginDto 用户登录
type LoginDto struct {
	Account  string `json:"account"`  // 手机或者用户名
	Password string `json:"password"` // 密码
}

// QueryUserMenuDtoResp 用户菜单
type QueryUserMenuDtoResp struct {
	Id       int64         `json:"id"`
	UserName string        `json:"name"`
	Avatar   string        `json:"avatar"`
	Menus    []UserMenuDto `json:"sysMenu"` // 菜单类型(1：目录   2：菜单   3：按钮)
	ApiUrls  []string      `json:"btnMenu"` // 接口权限
}
type UserMenuDto struct {
	Id       int64  `json:"id"`       // 主键
	MenuName string `json:"name"`     // 菜单名称
	ParentId int64  `json:"parentId"` // 父ID
	MenuUrl  string `json:"path"`     // 路由路径
	MenuIcon string `json:"icon"`     // 菜单图标
}

// QueryUserRoleListDto 查询用户已分配角色请求参数
type QueryUserRoleListDto struct {
	PageNo   int   `json:"pageNo"`   // 第几页
	PageSize int   `json:"pageSize"` // 每页的数量
	UserId   int64 `json:"userId" `
}
type QueryUserRoleListDataDtoResp struct {
	RoleList []QueryRoleListDtoResp `json:"sysRoleList"` // 角色列表
	RoleIds  []int64                `json:"userRoleIds"` // 角色Id
}

// UpdateUserRoleDto 更新用户角色参数
type UpdateUserRoleDto struct {
	UserId  int64   `json:"userId" binding:"required"`
	RoleIds []int64 `json:"roleIds" binding:"required"`
}
