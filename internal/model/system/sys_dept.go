package system

import (
	"time"
)

// SysDept 部门表
type SysDept struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                           // 部门id
	ParentId   int64     `gorm:"column:parent_id;default:0;NOT NULL" json:"parent_id"`                     // 父部门id
	Ancestors  string    `gorm:"column:ancestors;NOT NULL" json:"ancestors"`                               // 祖级列表
	DeptName   string    `gorm:"column:dept_name;NOT NULL" json:"dept_name"`                               // 部门名称
	Sort       int       `gorm:"column:sort;default:0;NOT NULL" json:"sort"`                               // 显示顺序
	Leader     string    `gorm:"column:leader;NOT NULL" json:"leader"`                                     // 负责人
	Phone      string    `gorm:"column:phone;NOT NULL" json:"phone"`                                       // 联系电话
	Email      string    `gorm:"column:email;NOT NULL" json:"email"`                                       // 邮箱
	Status     int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 部门状态（0：停用，1:正常）
	DelFlag    int       `gorm:"column:del_flag;default:1;NOT NULL" json:"del_flag"`                       // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysDept) TableName() string {
	return "sys_dept"
}
