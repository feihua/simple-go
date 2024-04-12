package models

import "time"

type OperationLog struct {
	Id            int64     `json:"id"`            //编号
	UserName      string    `json:"userName"`      //用户名
	Operation     string    `json:"operation"`     //用户操作
	Method        string    `json:"method"`        //请求方法
	Params        string    `json:"params"`        //请求参数
	OperationTime int64     `json:"operationTime"` //执行时长(毫秒)
	Ip            string    `json:"ip"`            //IP地址
	CreateTime    time.Time `json:"createTime"`    //创建时间
}

func (model OperationLog) TableName() string {
	return "sys_operation_log"
}
