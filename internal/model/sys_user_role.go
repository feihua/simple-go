package model

import (
	"time"
)

const TableNameUserRole = "sys_user_role"

// UserRole 角色用户关联表
type UserRole struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                         // 主键
	UserId     int64      `gorm:"column:user_id;not null;comment:用户Id" json:"userId"`                                   // 用户Id
	RoleId     int64      `gorm:"column:role_id;not null;comment:角色Id" json:"roleId"`                                   // 角色Id
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`                                    // 修改时间
}

// TableName UserRole's table name
func (*UserRole) TableName() string {
	return TableNameUserRole
}
