package models

import "time"

type RoleMenu struct {
	Id         int64     `json:"id"`         //主键
	RoleId     int64     `json:"roleId"`     //角色ID
	MenuId     int64     `json:"menuId"`     //菜单ID
	StatusId   int8      `json:"statusId"`   //状态(1:正常，0:禁用)
	Sort       int32     `json:"sort"`       //排序
	CreateTime time.Time `json:"createTime"` //创建时间
}

func (model RoleMenu) TableName() string {
	return "sys_role_menu"
}
