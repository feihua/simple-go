package system

import "time"

// OperateLog 操作日志记录
type OperateLog struct {
	Id              int64     `gorm:"column:id" json:"id"`                            // 日志主键
	Title           string    `gorm:"column:title" json:"title"`                      // 模块标题
	BusinessType    int32     `gorm:"column:business_type" json:"businessType"`       // 业务类型（0其它 1新增 2修改 3删除）
	Method          string    `gorm:"column:method" json:"method"`                    // 方法名称
	RequestMethod   string    `gorm:"column:request_method" json:"requestMethod"`     // 请求方式
	OperatorType    int32     `gorm:"column:operator_type" json:"operatorType"`       // 操作类别（0其它 1后台用户 2手机端用户）
	OperateName     string    `gorm:"column:operate_name" json:"operateName"`         // 操作人员
	DeptName        string    `gorm:"column:dept_name" json:"deptName"`               // 部门名称
	OperateUrl      string    `gorm:"column:operate_url" json:"operateUrl"`           // 请求URL
	OperateIp       string    `gorm:"column:operate_ip" json:"operateIp"`             // 主机地址
	OperateLocation string    `gorm:"column:operate_location" json:"operateLocation"` // 操作地点
	OperateParam    string    `gorm:"column:operate_param" json:"operateParam"`       // 请求参数
	JsonResult      string    `gorm:"column:json_result" json:"jsonResult"`           // 返回参数
	Platform        string    `gorm:"column:platform" json:"platform"`                // 平台信息
	Browser         string    `gorm:"column:browser" json:"browser"`                  // 浏览器类型
	Version         string    `gorm:"column:version" json:"version"`                  // 浏览器版本
	Os              string    `gorm:"column:os" json:"os"`                            // 操作系统
	Arch            string    `gorm:"column:arch" json:"arch"`                        // 体系结构信息
	Engine          string    `gorm:"column:engine" json:"engine"`                    // 渲染引擎信息
	EngineDetails   string    `gorm:"column:engine_details" json:"engineDetails"`     // 渲染引擎详细信息
	Extra           string    `gorm:"column:extra" json:"extra"`                      // 其他信息（可选）
	Status          int32     `gorm:"column:status" json:"status"`                    // 操作状态(0:异常,正常)
	ErrorMsg        string    `gorm:"column:error_msg" json:"errorMsg"`               // 错误消息
	OperateTime     time.Time `gorm:"column:operate_time" json:"operateTime"`         // 操作时间
	CostTime        int64     `gorm:"column:cost_time" json:"costTime"`               // 消耗时间
}

func (model OperateLog) TableName() string {
	return "sys_operate_log"
}
