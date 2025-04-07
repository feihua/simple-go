package model

import "time"

type Dict struct {
	Id          int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"` // 编号
	Value       string     `gorm:"column:value;not null;comment:数据值" json:"value"`               // 数据值
	Label       string     `gorm:"column:label;not null;comment:标签名" json:"label"`               // 标签名
	Type        string     `gorm:"column:type;not null;comment:类型" json:"type"`                  // 类型
	Description string     `gorm:"column:description;not null;comment:描述" json:"description"`    // 描述
	Sort        int32      `gorm:"column:sort;not null;comment:排序" json:"sort"`
	Remark      *string    `gorm:"column:remark;comment:备注" json:"remark"`
	CreateTime  time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime  *time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`
}

func (*Dict) TableName() string {
	return "sys_dict"
}
