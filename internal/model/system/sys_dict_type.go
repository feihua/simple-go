package system

import (
	"time"
)

// SysDictType 字典类型表
type SysDictType struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 字典主键
	DictName   string    `gorm:"column:dict_name;NOT NULL" json:"dict_name"`                               // 字典名称
	DictType   string    `gorm:"column:dict_type;NOT NULL" json:"dict_type"`                               // 字典类型
	Status     int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 状态（0：停用，1:正常）
	Remark     string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 备注
	CreateBy   string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysDictType) TableName() string {
	return "sys_dict_type"
}
