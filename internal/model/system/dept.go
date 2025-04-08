package system

import "time"

// Dept 部门
type Dept struct {
	Id         int64     `gorm:"column:id" json:"id"`                  // 部门id
	ParentId   int64     `gorm:"column:parent_id" json:"parentId"`     // 父部门id
	Ancestors  string    `gorm:"column:ancestors" json:"ancestors"`    // 祖级列表
	DeptName   string    `gorm:"column:dept_name" json:"deptName"`     // 部门名称
	Sort       int32     `gorm:"column:sort" json:"sort"`              // 显示顺序
	Leader     string    `gorm:"column:leader" json:"leader"`          // 负责人
	Phone      string    `gorm:"column:phone" json:"phone"`            // 联系电话
	Email      string    `gorm:"column:email" json:"email"`            // 邮箱
	Status     int32     `gorm:"column:status" json:"status"`          // 部门状态（0：停用，1:正常）
	DelFlag    int32     `gorm:"column:del_flag" json:"delFlag"`       // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`     // 创建者
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`     // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

func (model Dept) TableName() string {
	return "sys_dept"
}
