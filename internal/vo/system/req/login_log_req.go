package system

// DeleteLoginLogReqVo 删除系统访问记录请求参数
type DeleteLoginLogReqVo struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// QueryLoginLogDetailReqVo 查询系统访问记录详情请求参数
type QueryLoginLogDetailReqVo struct {
	Id int64 `json:"id" binding:"required"` // id
}

// QueryLoginLogListReqVo 查询系统访问记录列表请求参数
type QueryLoginLogListReqVo struct {
	PageNo        int    `json:"pageNo" default:"1" binding:"required"`    // 第几页
	PageSize      int    `json:"pageSize" default:"10" binding:"required"` // 每页的数量
	LoginName     string `json:"loginName" binding:"required"`             // 登录账号
	Ipaddr        string `json:"ipaddr" binding:"required"`                // 登录IP地址
	LoginLocation string `json:"loginLocation" binding:"required"`         // 登录地点
	Platform      string `json:"platform" binding:"required"`              // 平台信息
	Browser       string `json:"browser" binding:"required"`               // 浏览器类型
	Version       string `json:"version" binding:"required"`               // 浏览器版本
	Os            string `json:"os" binding:"required"`                    // 操作系统
	Status        int32  `json:"status" binding:"required"`                // 登录状态(0:失败,1:成功)

}
