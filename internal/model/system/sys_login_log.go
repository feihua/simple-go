package system

import (
	"time"
)

// 系统访问记录
type SysLoginLog struct {
	Id            int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                         // 访问ID
	LoginName     string    `gorm:"column:login_name;NOT NULL" json:"login_name"`                           // 登录账号
	Ipaddr        string    `gorm:"column:ipaddr;NOT NULL" json:"ipaddr"`                                   // 登录IP地址
	LoginLocation string    `gorm:"column:login_location;NOT NULL" json:"login_location"`                   // 登录地点
	Platform      string    `gorm:"column:platform;NOT NULL" json:"platform"`                               // 平台信息
	Browser       string    `gorm:"column:browser;NOT NULL" json:"browser"`                                 // 浏览器类型
	Version       string    `gorm:"column:version;NOT NULL" json:"version"`                                 // 浏览器版本
	Os            string    `gorm:"column:os;NOT NULL" json:"os"`                                           // 操作系统
	Arch          string    `gorm:"column:arch;NOT NULL" json:"arch"`                                       // 体系结构信息
	Engine        string    `gorm:"column:engine;NOT NULL" json:"engine"`                                   // 渲染引擎信息
	EngineDetails string    `gorm:"column:engine_details;NOT NULL" json:"engine_details"`                   // 渲染引擎详细信息
	Extra         string    `gorm:"column:extra;NOT NULL" json:"extra"`                                     // 其他信息（可选）
	Status        int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                         // 登录状态(0:失败,1:成功)
	Msg           string    `gorm:"column:msg;NOT NULL" json:"msg"`                                         // 提示消息
	LoginTime     time.Time `gorm:"column:login_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"login_time"` // 访问时间
}

func (m *SysLoginLog) TableName() string {
	return "sys_login_log"
}
