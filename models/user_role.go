package models

import "time"

type UserRole struct {
	Id         int64     `json:"id"`         //主键
	UserId     int64     `json:"userId"`     //用户ID
	RoleId     int64     `json:"roleId"`     //角色ID
	StatusId   int8      `json:"statusId"`   //状态(1:正常，0:禁用)
	Sort       int32     `json:"sort"`       //排序
	CreateTime time.Time `json:"createTime"` //创建时间
}

func (model UserRole) TableName() string {
	return "sys_user_role"
}
