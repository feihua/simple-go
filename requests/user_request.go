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

// LoginRequest 登陆入参
type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginRequest 登陆入参
type UserRoleRequest struct {
	UserId int64 `json:"userId"`
	Roleid int64 `json:"roleId" `
}
