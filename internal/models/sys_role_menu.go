package models

import (
	"time"
)

const TableNameRoleMenu = "sys_role_menu"

// RoleMenu 菜单角色关联表
type RoleMenu struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                         // 主键
	RoleId     int64      `gorm:"column:role_id;not null;comment:角色Id" json:"roleId"`                                   // 角色Id
	MenuId     int64      `gorm:"column:menu_id;not null;comment:菜单Id" json:"menuId"`                                   // 菜单Id
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`                                    // 修改时间
}

// TableName RoleMenu's table name
func (*RoleMenu) TableName() string {
	return TableNameRoleMenu
}
