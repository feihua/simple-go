package models

import "time"

type Menu struct {
	Id         int64     `json:"id"`         //主键
	MenuName   string    `json:"menuName"`   //菜单名称
	MenuType   int8      `json:"menuType"`   //菜单类型(1：目录   2：菜单   3：按钮)
	StatusId   int8      `json:"statusId"`   //状态(1:正常，0:禁用)
	Sort       int32     `json:"sort"`       //排序
	ParentId   int64     `json:"parentId"`   //父ID
	MenuUrl    string    `json:"menuUrl"`    //路由路径
	ApiUrl     string    `json:"apiUrl"`     //接口URL
	MenuIcon   string    `json:"menuIcon"`   //菜单图标
	Remark     string    `json:"remark"`     //备注
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateTime time.Time `json:"updateTime"` //修改时间
}

func (model Menu) TableName() string {
	return "sys_menu"
}
