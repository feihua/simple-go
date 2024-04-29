package dto

type UserDto struct {
	Id       int64  `json:"id"`       //主键
	Mobile   string `json:"mobile"`   //手机
	UserName string `json:"userName"` //姓名
	Password string `json:"password"` //密码
	StatusId int8   `json:"statusId"` //状态(1:正常，0:禁用)
	Sort     int32  `json:"sort"`     //排序
	Remark   string `json:"remark"`   //备注
}

type UserLoginDto struct {
	Account  string `json:"account"`  //手机或者用户名
	Password string `json:"password"` //密码
}

// LoginDtoResp 登录响应
type LoginDtoResp struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

// QueryUserMenuDtoResp 用户菜单
type QueryUserMenuDtoResp struct {
	Id       int64         `json:"id"`
	UserName string        `json:"userName"`
	Avatar   string        `json:"avatar"`
	Menus    []UserMenuDto `json:"Menus"`   //菜单类型(1：目录   2：菜单   3：按钮)
	ApiUrls  []string      `json:"apiUrls"` //接口权限
}
type UserMenuDto struct {
	Id       int64  `json:"id"`       //主键
	MenuName string `json:"menuName"` //菜单名称
	Sort     int32  `json:"sort"`     //排序
	ParentId int64  `json:"parentId"` //父ID
	MenuUrl  string `json:"menuUrl"`  //路由路径
	MenuIcon string `json:"menuIcon"` //菜单图标
}

// UpdateUserRoleDtoRequest 更新用户角色参数
type UpdateUserRoleDtoRequest struct {
	UserId int64   `json:"userId"`
	RoleId []int64 `json:"roleId" `
}

// QueryUserListDto 查询用户列表
type QueryUserListDto struct {
	Mobile   string `json:"mobile"`   //手机
	UserName string `json:"userName"` //姓名
	StatusId int8   `json:"statusId"` //状态(1:正常，0:禁用)
	PageNo   int    `json:"pageNo"`   //第几页
	PageSize int    `json:"pageSize"` //每页的数量
}
