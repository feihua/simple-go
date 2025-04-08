package system

import (
	"time"
)

// SysOperateLog 操作日志记录
type SysOperateLog struct {
	Id              int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                             // 日志主键
	Title           string    `gorm:"column:title;NOT NULL" json:"title"`                                         // 模块标题
	BusinessType    int       `gorm:"column:business_type;default:0;NOT NULL" json:"business_type"`               // 业务类型（0其它 1新增 2修改 3删除）
	Method          string    `gorm:"column:method;NOT NULL" json:"method"`                                       // 方法名称
	RequestMethod   string    `gorm:"column:request_method;NOT NULL" json:"request_method"`                       // 请求方式
	OperatorType    int       `gorm:"column:operator_type;default:0;NOT NULL" json:"operator_type"`               // 操作类别（0其它 1后台用户 2手机端用户）
	OperateName     string    `gorm:"column:operate_name;NOT NULL" json:"operate_name"`                           // 操作人员
	DeptName        string    `gorm:"column:dept_name;NOT NULL" json:"dept_name"`                                 // 部门名称
	OperateUrl      string    `gorm:"column:operate_url;NOT NULL" json:"operate_url"`                             // 请求URL
	OperateIp       string    `gorm:"column:operate_ip;NOT NULL" json:"operate_ip"`                               // 主机地址
	OperateLocation string    `gorm:"column:operate_location;NOT NULL" json:"operate_location"`                   // 操作地点
	OperateParam    string    `gorm:"column:operate_param;NOT NULL" json:"operate_param"`                         // 请求参数
	JsonResult      string    `gorm:"column:json_result;NOT NULL" json:"json_result"`                             // 返回参数
	Platform        string    `gorm:"column:platform;NOT NULL" json:"platform"`                                   // 平台信息
	Browser         string    `gorm:"column:browser;NOT NULL" json:"browser"`                                     // 浏览器类型
	Version         string    `gorm:"column:version;NOT NULL" json:"version"`                                     // 浏览器版本
	Os              string    `gorm:"column:os;NOT NULL" json:"os"`                                               // 操作系统
	Arch            string    `gorm:"column:arch;NOT NULL" json:"arch"`                                           // 体系结构信息
	Engine          string    `gorm:"column:engine;NOT NULL" json:"engine"`                                       // 渲染引擎信息
	EngineDetails   string    `gorm:"column:engine_details;NOT NULL" json:"engine_details"`                       // 渲染引擎详细信息
	Extra           string    `gorm:"column:extra;NOT NULL" json:"extra"`                                         // 其他信息（可选）
	Status          int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                             // 操作状态(0:异常,正常)
	ErrorMsg        string    `gorm:"column:error_msg;NOT NULL" json:"error_msg"`                                 // 错误消息
	OperateTime     time.Time `gorm:"column:operate_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"operate_time"` // 操作时间
	CostTime        int64     `gorm:"column:cost_time;default:0;NOT NULL" json:"cost_time"`                       // 消耗时间
}

func (m *SysOperateLog) TableName() string {
	return "sys_operate_log"
}
