package requests

type UserRequest struct {
	Id       int64  `json:"id"`       //主键
	Mobile   string `json:"mobile"`   //手机
	UserName string `json:"userName"` //姓名
	Password string `json:"password"` //密码
	StatusId int8   `json:"statusId"` //状态(1:正常，0:禁用)
	Sort     int32  `json:"sort"`     //排序
	Remark   string `json:"remark"`   //备注
}

// RegisterRequest 注册入参
type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginRequest 登陆参数
type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserRoleRequest 更新用户角色参数
type UpdateUserRoleRequest struct {
	UserId int64   `json:"userId"`
	RoleId []int64 `json:"roleId" `
}

// DeleteUserRequest 删除用户
type DeleteUserRequest struct {
	Ids []int64 `json:"ids"` //编号
}
