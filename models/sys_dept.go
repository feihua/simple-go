package models

import "time"

type Dept struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"` // 编号
	DeptName   string     `gorm:"column:dept_name;not null;comment:部门名称" json:"deptName"`
	ParentId   int64      `gorm:"column:parent_id;not null;comment:上级部门Id，一级部门为0" json:"parentId"`
	Sort       int32      `gorm:"column:sort;not null;comment:排序" json:"sort"`
	Remark     *string    `gorm:"column:remark;comment:备注" json:"remark"`
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`
}

func (*Dept) TableName() string {
	return "sys_dept"
}
