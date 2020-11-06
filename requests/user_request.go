package requests

import "time"

type UserRequest struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Name           string    `json:"name" xorm:"not null comment('用户名') VARCHAR(50) 'name'"`
	NickName       string    `json:"nick_name" xorm:"comment('昵称') VARCHAR(150) 'nick_name'"`
	Avatar         string    `json:"avatar" xorm:"comment('头像') VARCHAR(150) 'avatar'"`
	Password       string    `json:"password" xorm:"comment('密码') VARCHAR(100) 'password'"`
	Salt           string    `json:"salt" xorm:"comment('加密盐') VARCHAR(40) 'salt'"`
	Email          string    `json:"email" xorm:"comment('邮箱') VARCHAR(100) 'email'"`
	Mobile         string    `json:"mobile" xorm:"comment('手机号') VARCHAR(100) 'mobile'"`
	Status         int64     `json:"status" xorm:"comment('状态  0：禁用   1：正常') TINYINT(4) 'status'"`
	DeptId         int64     `json:"dept_id" xorm:"comment('机构ID') BIGINT(20) 'dept_id'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
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
