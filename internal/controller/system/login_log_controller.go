package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/login_log"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// LoginLogController 系统访问记录相关
type LoginLogController struct {
	Service login_log.LoginLogService
}

func NewLoginLogController(Service login_log.LoginLogService) *LoginLogController {
	return &LoginLogController{Service: Service}
}

// CreateLoginLog 添加系统访问记录
func (r LoginLogController) CreateLoginLog(c *gin.Context) {

	req := b.AddLoginLogReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.AddLoginLogDto{
		Id:            req.Id,            // 访问ID
		LoginName:     req.LoginName,     // 登录账号
		Ipaddr:        req.Ipaddr,        // 登录IP地址
		LoginLocation: req.LoginLocation, // 登录地点
		Platform:      req.Platform,      // 平台信息
		Browser:       req.Browser,       // 浏览器类型
		Version:       req.Version,       // 浏览器版本
		Os:            req.Os,            // 操作系统
		Arch:          req.Arch,          // 体系结构信息
		Engine:        req.Engine,        // 渲染引擎信息
		EngineDetails: req.EngineDetails, // 渲染引擎详细信息
		Extra:         req.Extra,         // 其他信息（可选）
		Status:        req.Status,        // 登录状态(0:失败,1:成功)
		Msg:           req.Msg,           // 提示消息
		LoginTime:     req.LoginTime,     // 访问时间
	}

	err = r.Service.CreateLoginLog(item)
	if err != nil {
		result.FailWithMsg(c, result.LoginLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteLoginLogByIds 删除系统访问记录
func (r LoginLogController) DeleteLoginLogByIds(c *gin.Context) {

	req := b.DeleteLoginLogReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteLoginLogByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.LoginLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateLoginLog 更新系统访问记录
func (r LoginLogController) UpdateLoginLog(c *gin.Context) {

	req := b.UpdateLoginLogReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateLoginLogDto{
		Id:            req.Id,            // 访问ID
		LoginName:     req.LoginName,     // 登录账号
		Ipaddr:        req.Ipaddr,        // 登录IP地址
		LoginLocation: req.LoginLocation, // 登录地点
		Platform:      req.Platform,      // 平台信息
		Browser:       req.Browser,       // 浏览器类型
		Version:       req.Version,       // 浏览器版本
		Os:            req.Os,            // 操作系统
		Arch:          req.Arch,          // 体系结构信息
		Engine:        req.Engine,        // 渲染引擎信息
		EngineDetails: req.EngineDetails, // 渲染引擎详细信息
		Extra:         req.Extra,         // 其他信息（可选）
		Status:        req.Status,        // 登录状态(0:失败,1:成功)
		Msg:           req.Msg,           // 提示消息
		LoginTime:     req.LoginTime,     // 访问时间
	}
	err = r.Service.UpdateLoginLog(item)
	if err != nil {
		result.FailWithMsg(c, result.LoginLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateLoginLogStatus 更新系统访问记录状态
func (r LoginLogController) UpdateLoginLogStatus(c *gin.Context) {

	req := b.UpdateLoginLogStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateLoginLogStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
	}
	err = r.Service.UpdateLoginLogStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.LoginLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryLoginLogDetail 查询系统访问记录详情
func (r LoginLogController) QueryLoginLogDetail(c *gin.Context) {
	req := b.QueryLoginLogDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryLoginLogDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryLoginLogDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.LoginLogError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryLoginLogList 查询系统访问记录列表
func (r LoginLogController) QueryLoginLogList(c *gin.Context) {
	req := b.QueryLoginLogListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryLoginLogListDto{
		PageNo:        req.PageNo,
		PageSize:      req.PageSize,
		LoginName:     req.LoginName,     // 登录账号
		Ipaddr:        req.Ipaddr,        // 登录IP地址
		LoginLocation: req.LoginLocation, // 登录地点
		Platform:      req.Platform,      // 平台信息
		Browser:       req.Browser,       // 浏览器类型
		Version:       req.Version,       // 浏览器版本
		Os:            req.Os,            // 操作系统
		Arch:          req.Arch,          // 体系结构信息
		Engine:        req.Engine,        // 渲染引擎信息
		EngineDetails: req.EngineDetails, // 渲染引擎详细信息
		Extra:         req.Extra,         // 其他信息（可选）
		Status:        req.Status,        // 登录状态(0:失败,1:成功)
		Msg:           req.Msg,           // 提示消息
	}
	list, total := r.Service.QueryLoginLogList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
