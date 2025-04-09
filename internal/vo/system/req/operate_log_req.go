package system

// DeleteOperateLogReqVo 删除操作日志记录请求参数
type DeleteOperateLogReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// QueryOperateLogDetailReqVo 查询操作日志记录详情请求参数
type QueryOperateLogDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryOperateLogListReqVo 查询操作日志记录列表请求参数
type QueryOperateLogListReqVo struct {
	PageNo          int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize        int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	Id              int64  `json:"id" binding:"required"`                    // 日志主键
	Title           string `json:"title" binding:"required"`                 // 模块标题
	BusinessType    int32  `json:"businessType" binding:"required"`          // 业务类型（0其它 1新增 2修改 3删除）
	Method          string `json:"method" binding:"required"`                // 方法名称
	RequestMethod   string `json:"requestMethod" binding:"required"`         // 请求方式
	OperatorType    int32  `json:"operatorType" binding:"required"`          // 操作类别（0其它 1后台用户 2手机端用户）
	OperateName     string `json:"operateName" binding:"required"`           // 操作人员
	DeptName        string `json:"deptName" binding:"required"`              // 部门名称
	OperateUrl      string `json:"operateUrl" binding:"required"`            // 请求URL
	OperateIp       string `json:"operateIp" binding:"required"`             // 主机地址
	OperateLocation string `json:"operateLocation" binding:"required"`       // 操作地点
	Platform        string `json:"platform" binding:"required"`              // 平台信息
	Browser         string `json:"browser" binding:"required"`               // 浏览器类型
	Version         string `json:"version" binding:"required"`               // 浏览器版本
	Os              string `json:"os" binding:"required"`                    // 操作系统
	Status          int32  `json:"status" binding:"required"`                // 操作状态(0:异常,正常)

}
