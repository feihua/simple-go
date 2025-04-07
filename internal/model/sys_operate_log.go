package model

import (
	"time"
)

const TableNameOperateLog = "sys_operate_log"

// OperateLog 系统操作日志
type OperateLog struct {
	Id             int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                      // 编号
	UserName       string     `gorm:"column:user_name;not null;comment:用户名" json:"userName"`                             // 用户名
	Operation      string     `gorm:"column:operation;not null;comment:用户操作" json:"operation"`                           // 用户操作
	Method         string     `gorm:"column:method;not null;comment:请求方法" json:"method"`                                 // 请求方法
	RequestParams  string     `gorm:"column:request_params;not null;comment:请求参数" json:"requestParams"`                  // 请求参数
	ResponseParams *string    `gorm:"column:response_params;comment:响应参数" json:"responseParams"`                         // 响应参数
	Time           int64      `gorm:"column:time;not null;comment:执行时长(毫秒)" json:"time"`                                 // 执行时长(毫秒)
	Ip             *string    `gorm:"column:ip;comment:Ip地址" json:"ip"`                                                  // IP地址
	OperationTime  *time.Time `gorm:"column:operation_time;default:CURRENT_TIMESTAMP;comment:操作时间" json:"operationTime"` // 操作时间
}

// TableName OperateLog's table name
func (*OperateLog) TableName() string {
	return TableNameOperateLog
}
