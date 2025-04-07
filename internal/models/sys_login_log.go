package models

import (
	"time"
)

const TableNameLoginLog = "sys_login_log"

// LoginLog 系统登录日志
type LoginLog struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                       // 编号
	UserName  string    `gorm:"column:user_name;not null;comment:用户名" json:"userName"`                              // 用户名
	Status    string    `gorm:"column:status;not null;comment:登录状态" json:"status"`                                  // 登录状态
	Ip        string    `gorm:"column:ip;not null;comment:Ip地址" json:"ip"`                                          // IP地址
	LoginTime time.Time `gorm:"column:login_time;not null;default:CURRENT_TIMESTAMP;comment:登陆时间" json:"loginTime"` // 登陆时间
}

// TableName LoginLog's table name
func (*LoginLog) TableName() string {
	return TableNameLoginLog
}
