package models

import (
	"time"
)

const TableNameRole = "sys_role"

// Role 角色信息
type Role struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                         // 主键
	RoleName   string     `gorm:"column:role_name;not null;comment:名称" json:"roleName"`                                 // 名称
	StatusId   int32      `gorm:"column:status_id;not null;default:1;comment:状态(1:正常，0:禁用)" json:"statusId"`            // 状态(1:正常，0:禁用)
	Sort       int32      `gorm:"column:sort;not null;default:1;comment:排序" json:"sort"`                                // 排序
	Remark     string     `gorm:"column:remark;not null;comment:备注" json:"remark"`                                      // 备注
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`                                    // 修改时间
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
