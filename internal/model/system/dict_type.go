package system

import "time"

// DictType 字典类型
type DictType struct {
	Id         int64      `gorm:"column:id" json:"id"`                                            // 字典主键
	DictName   string     `gorm:"column:dict_name" json:"dictName"`                               // 字典名称
	DictType   string     `gorm:"column:dict_type" json:"dictType"`                               // 字典类型
	Status     int32      `gorm:"column:status" json:"status"`                                    // 状态（0：停用，1:正常）
	Remark     string     `gorm:"column:remark" json:"remark"`                                    // 备注
	CreateBy   string     `gorm:"column:create_by" json:"createBy"`                               // 创建者
	CreateTime time.Time  `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by" json:"updateBy"`                               // 更新者
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime"`                           // 更新时间
}

func (model DictType) TableName() string {
	return "sys_dict_type"
}
