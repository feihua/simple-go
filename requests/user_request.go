package requests

type UserRequest struct {
	ID       int64   `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

//RegisterRequest 注册入参
type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//LoginRequest 登陆入参
type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//LoginRequest 登陆入参
type UserRoleRequest struct {
	UserId int64 `json:"userId"`
	Roleid int64 `json:"roleId" `
}