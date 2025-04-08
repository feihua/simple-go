package system

import "time"

// AddLoginLogReqVo 添加系统访问记录请求参数
type AddLoginLogReqVo struct {
	Id            int64     `json:"id"`            //访问ID
	LoginName     string    `json:"loginName"`     //登录账号
	Ipaddr        string    `json:"ipaddr"`        //登录IP地址
	LoginLocation string    `json:"loginLocation"` //登录地点
	Platform      string    `json:"platform"`      //平台信息
	Browser       string    `json:"browser"`       //浏览器类型
	Version       string    `json:"version"`       //浏览器版本
	Os            string    `json:"os"`            //操作系统
	Arch          string    `json:"arch"`          //体系结构信息
	Engine        string    `json:"engine"`        //渲染引擎信息
	EngineDetails string    `json:"engineDetails"` //渲染引擎详细信息
	Extra         string    `json:"extra"`         //其他信息（可选）
	Status        int32     `json:"status"`        //登录状态(0:失败,1:成功)
	Msg           string    `json:"msg"`           //提示消息
	LoginTime     time.Time `json:"loginTime"`     //访问时间
}

// DeleteLoginLogReqVo 删除系统访问记录请求参数
type DeleteLoginLogReqVo struct {
	Ids []int64 `json:"ids"`
}

// UpdateLoginLogReqVo 修改系统访问记录请求参数
type UpdateLoginLogReqVo struct {
	Id            int64     `json:"id"`            //访问ID
	LoginName     string    `json:"loginName"`     //登录账号
	Ipaddr        string    `json:"ipaddr"`        //登录IP地址
	LoginLocation string    `json:"loginLocation"` //登录地点
	Platform      string    `json:"platform"`      //平台信息
	Browser       string    `json:"browser"`       //浏览器类型
	Version       string    `json:"version"`       //浏览器版本
	Os            string    `json:"os"`            //操作系统
	Arch          string    `json:"arch"`          //体系结构信息
	Engine        string    `json:"engine"`        //渲染引擎信息
	EngineDetails string    `json:"engineDetails"` //渲染引擎详细信息
	Extra         string    `json:"extra"`         //其他信息（可选）
	Status        int32     `json:"status"`        //登录状态(0:失败,1:成功)
	Msg           string    `json:"msg"`           //提示消息
	LoginTime     time.Time `json:"loginTime"`     //访问时间
}

// UpdateLoginLogStatusReqVo 修改系统访问记录状态请求参数
type UpdateLoginLogStatusReqVo struct {
	Ids    []int64 `json:"ids"`    //id
	Status int32   `json:"status"` //状态（0:关闭,1:正常 ）
}

// QueryLoginLogDetailReqVo 查询系统访问记录详情请求参数
type QueryLoginLogDetailReqVo struct {
	Id int64 `json:"id"` //id
}

// QueryLoginLogListReqVo 查询系统访问记录列表请求参数
type QueryLoginLogListReqVo struct {
	PageNo        int       `json:"pageNo" default:"1"`    //第几页
	PageSize      int       `json:"pageSize" default:"10"` //每页的数量
	Id            int64     `json:"id"`                    //访问ID
	LoginName     string    `json:"loginName"`             //登录账号
	Ipaddr        string    `json:"ipaddr"`                //登录IP地址
	LoginLocation string    `json:"loginLocation"`         //登录地点
	Platform      string    `json:"platform"`              //平台信息
	Browser       string    `json:"browser"`               //浏览器类型
	Version       string    `json:"version"`               //浏览器版本
	Os            string    `json:"os"`                    //操作系统
	Arch          string    `json:"arch"`                  //体系结构信息
	Engine        string    `json:"engine"`                //渲染引擎信息
	EngineDetails string    `json:"engineDetails"`         //渲染引擎详细信息
	Extra         string    `json:"extra"`                 //其他信息（可选）
	Status        int32     `json:"status"`                //登录状态(0:失败,1:成功)
	Msg           string    `json:"msg"`                   //提示消息
	LoginTime     time.Time `json:"loginTime"`             //访问时间
}
