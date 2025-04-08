package system

import "time"

// OperateLog 操作日志记录
type OperateLog struct {
	Id              int64     `json:"id"`              //日志主键
	Title           string    `json:"title"`           //模块标题
	BusinessType    int32     `json:"businessType"`    //业务类型（0其它 1新增 2修改 3删除）
	Method          string    `json:"method"`          //方法名称
	RequestMethod   string    `json:"requestMethod"`   //请求方式
	OperatorType    int32     `json:"operatorType"`    //操作类别（0其它 1后台用户 2手机端用户）
	OperateName     string    `json:"operateName"`     //操作人员
	DeptName        string    `json:"deptName"`        //部门名称
	OperateUrl      string    `json:"operateUrl"`      //请求URL
	OperateIp       string    `json:"operateIp"`       //主机地址
	OperateLocation string    `json:"operateLocation"` //操作地点
	OperateParam    string    `json:"operateParam"`    //请求参数
	JsonResult      string    `json:"jsonResult"`      //返回参数
	Platform        string    `json:"platform"`        //平台信息
	Browser         string    `json:"browser"`         //浏览器类型
	Version         string    `json:"version"`         //浏览器版本
	Os              string    `json:"os"`              //操作系统
	Arch            string    `json:"arch"`            //体系结构信息
	Engine          string    `json:"engine"`          //渲染引擎信息
	EngineDetails   string    `json:"engineDetails"`   //渲染引擎详细信息
	Extra           string    `json:"extra"`           //其他信息（可选）
	Status          int32     `json:"status"`          //操作状态(0:异常,正常)
	ErrorMsg        string    `json:"errorMsg"`        //错误消息
	OperateTime     time.Time `json:"operateTime"`     //操作时间
	CostTime        int64     `json:"costTime"`        //消耗时间
}
