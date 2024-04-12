package models

import "time"

type Dict struct {
	Id          int64     `json:"id"`          //编号
	Value       string    `json:"value"`       //数据值
	Label       string    `json:"label"`       //标签名
	Type        string    `json:"type"`        //类型
	Description string    `json:"description"` //描述
	Sort        float64   `json:"sort"`        //排序（升序）
	Remarks     string    `json:"remarks"`     //备注
	CreateTime  time.Time `json:"createTime"`  //创建时间
	UpdateTime  time.Time `json:"updateTime"`  //修改时间
}

func (model Dict) TableName() string {
	return "sys_dict"
}
