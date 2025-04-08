package system

import "time"

// DictData 字典数据
type DictData struct {
	Id         int64     `json:"id"`         //字典编码
	DictSort   int32     `json:"dictSort"`   //字典排序
	DictLabel  string    `json:"dictLabel"`  //字典标签
	DictValue  string    `json:"dictValue"`  //字典键值
	DictType   string    `json:"dictType"`   //字典类型
	CssClass   string    `json:"cssClass"`   //样式属性（其他样式扩展）
	ListClass  string    `json:"listClass"`  //表格回显样式
	IsDefault  string    `json:"isDefault"`  //是否默认（Y是 N否）
	Status     int32     `json:"status"`     //状态（0：停用，1:正常）
	Remark     string    `json:"remark"`     //备注
	CreateBy   string    `json:"createBy"`   //创建者
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateBy   string    `json:"updateBy"`   //更新者
	UpdateTime time.Time `json:"updateTime"` //更新时间
}
