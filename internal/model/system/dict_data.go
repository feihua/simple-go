package system

import "time"

// DictData 字典数据
type DictData struct {
	Id         int64     `gorm:"column:id" json:"id"`                  // 字典编码
	DictSort   int32     `gorm:"column:dict_sort" json:"dictSort"`     // 字典排序
	DictLabel  string    `gorm:"column:dict_label" json:"dictLabel"`   // 字典标签
	DictValue  string    `gorm:"column:dict_value" json:"dictValue"`   // 字典键值
	DictType   string    `gorm:"column:dict_type" json:"dictType"`     // 字典类型
	CssClass   string    `gorm:"column:css_class" json:"cssClass"`     // 样式属性（其他样式扩展）
	ListClass  string    `gorm:"column:list_class" json:"listClass"`   // 表格回显样式
	IsDefault  string    `gorm:"column:is_default" json:"isDefault"`   // 是否默认（Y是 N否）
	Status     int32     `gorm:"column:status" json:"status"`          // 状态（0：停用，1:正常）
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`     // 创建者
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`     // 更新者
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

func (model DictData) TableName() string {
	return "sys_dict_data"
}
