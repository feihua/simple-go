package models

import "time"

type User struct {
	Id         int64     `json:"id"`         //主键
	Mobile     string    `json:"mobile"`     //手机
	UserName   string    `json:"userName"`   //姓名
	Password   string    `json:"password"`   //密码
	StatusId   int8      `json:"statusId"`   //状态(1:正常，0:禁用)
	Sort       int32     `json:"sort"`       //排序
	Remark     string    `json:"remark"`     //备注
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateTime time.Time `json:"updateTime"` //修改时间
}

func (model User) TableName() string {
	return "sys_user"
}
