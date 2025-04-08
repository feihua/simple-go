package system

import "time"

// LoginLog 系统访问记录
type LoginLog struct {
	Id            int64     `gorm:"column:id" json:"id"`                        // 访问ID
	LoginName     string    `gorm:"column:login_name" json:"loginName"`         // 登录账号
	Ipaddr        string    `gorm:"column:ipaddr" json:"ipaddr"`                // 登录IP地址
	LoginLocation string    `gorm:"column:login_location" json:"loginLocation"` // 登录地点
	Platform      string    `gorm:"column:platform" json:"platform"`            // 平台信息
	Browser       string    `gorm:"column:browser" json:"browser"`              // 浏览器类型
	Version       string    `gorm:"column:version" json:"version"`              // 浏览器版本
	Os            string    `gorm:"column:os" json:"os"`                        // 操作系统
	Arch          string    `gorm:"column:arch" json:"arch"`                    // 体系结构信息
	Engine        string    `gorm:"column:engine" json:"engine"`                // 渲染引擎信息
	EngineDetails string    `gorm:"column:engine_details" json:"engineDetails"` // 渲染引擎详细信息
	Extra         string    `gorm:"column:extra" json:"extra"`                  // 其他信息（可选）
	Status        int32     `gorm:"column:status" json:"status"`                // 登录状态(0:失败,1:成功)
	Msg           string    `gorm:"column:msg" json:"msg"`                      // 提示消息
	LoginTime     time.Time `gorm:"column:login_time" json:"loginTime"`         // 访问时间
}

func (model LoginLog) TableName() string {
	return "sys_login_log"
}
