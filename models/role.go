package models

import "time"

type Role struct {
	Id         int64     `json:"id"`         //主键
	RoleName   string    `json:"roleName"`   //名称
	StatusId   int8      `json:"statusId"`   //状态(1:正常，0:禁用)
	Sort       int32     `json:"sort"`       //排序
	Remark     string    `json:"remark"`     //备注
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateTime time.Time `json:"updateTime"` //修改时间
}

func (model Role) TableName() string {
	return "sys_role"
}
