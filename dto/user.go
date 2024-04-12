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
