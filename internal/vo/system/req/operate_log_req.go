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
	Title           string `json:"title" `                                   // 模块标题
	BusinessType    int32  `json:"businessType" `                            // 业务类型（0其它 1新增 2修改 3删除）
	Method          string `json:"method" `                                  // 方法名称
	RequestMethod   string `json:"requestMethod" `                           // 请求方式
	OperatorType    int32  `json:"operatorType" `                            // 操作类别（0其它 1后台用户 2手机端用户）
	OperateName     string `json:"operateName" `                             // 操作人员
	DeptName        string `json:"deptName" `                                // 部门名称
	OperateUrl      string `json:"operateUrl" `                              // 请求URL
	OperateIp       string `json:"operateIp" `                               // 主机地址
	OperateLocation string `json:"operateLocation" `                         // 操作地点
	Platform        string `json:"platform" `                                // 平台信息
	Browser         string `json:"browser" `                                 // 浏览器类型
	Version         string `json:"version" `                                 // 浏览器版本
	Os              string `json:"os" `                                      // 操作系统
	Status          int32  `json:"status" `                                  // 操作状态(0:异常,正常)

}
