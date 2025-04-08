package system

import "time"

// DictType 字典类型
type DictType struct {
	Id         int64     `json:"id"`         //字典主键
	DictName   string    `json:"dictName"`   //字典名称
	DictType   string    `json:"dictType"`   //字典类型
	Status     int32     `json:"status"`     //状态（0：停用，1:正常）
	Remark     string    `json:"remark"`     //备注
	CreateBy   string    `json:"createBy"`   //创建者
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateBy   string    `json:"updateBy"`   //更新者
	UpdateTime time.Time `json:"updateTime"` //更新时间
}
