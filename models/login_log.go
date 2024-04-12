package models

import "time"

type LoginLog struct {
	Id         int64     `json:"id"`         //编号
	UserName   string    `json:"userName"`   //用户名
	Ip         string    `json:"ip"`         //IP地址
	CreateBy   string    `json:"createBy"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
}

func (model LoginLog) TableName() string {
	return "sys_login_log"
}
