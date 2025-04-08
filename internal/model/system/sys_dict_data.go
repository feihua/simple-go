package system

import (
	"time"
)

// SysDictData 字典数据表
type SysDictData struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 字典编码
	DictSort   int       `gorm:"column:dict_sort;default:0;NOT NULL" json:"dict_sort"`                     // 字典排序
	DictLabel  string    `gorm:"column:dict_label;NOT NULL" json:"dict_label"`                             // 字典标签
	DictValue  string    `gorm:"column:dict_value;NOT NULL" json:"dict_value"`                             // 字典键值
	DictType   string    `gorm:"column:dict_type;NOT NULL" json:"dict_type"`                               // 字典类型
	CssClass   string    `gorm:"column:css_class;NOT NULL" json:"css_class"`                               // 样式属性（其他样式扩展）
	ListClass  string    `gorm:"column:list_class;NOT NULL" json:"list_class"`                             // 表格回显样式
	IsDefault  string    `gorm:"column:is_default;default:N;NOT NULL" json:"is_default"`                   // 是否默认（Y是 N否）
	Status     int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 状态（0：停用，1:正常）
	Remark     string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 备注
	CreateBy   string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysDictData) TableName() string {
	return "sys_dict_data"
}
