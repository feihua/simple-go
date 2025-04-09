package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/operate_log"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// OperateLogController 操作日志记录相关
type OperateLogController struct {
	Service operate_log.OperateLogService
}

func NewOperateLogController(Service operate_log.OperateLogService) *OperateLogController {
	return &OperateLogController{Service: Service}
}

// CreateOperateLog 添加操作日志记录
func (r OperateLogController) CreateOperateLog(c *gin.Context) {

	req := b.AddOperateLogReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.AddOperateLogDto{
		Id:              req.Id,              // 日志主键
		Title:           req.Title,           // 模块标题
		BusinessType:    req.BusinessType,    // 业务类型（0其它 1新增 2修改 3删除）
		Method:          req.Method,          // 方法名称
		RequestMethod:   req.RequestMethod,   // 请求方式
		OperatorType:    req.OperatorType,    // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     req.OperateName,     // 操作人员
		DeptName:        req.DeptName,        // 部门名称
		OperateUrl:      req.OperateUrl,      // 请求URL
		OperateIp:       req.OperateIp,       // 主机地址
		OperateLocation: req.OperateLocation, // 操作地点
		OperateParam:    req.OperateParam,    // 请求参数
		JsonResult:      req.JsonResult,      // 返回参数
		Platform:        req.Platform,        // 平台信息
		Browser:         req.Browser,         // 浏览器类型
		Version:         req.Version,         // 浏览器版本
		Os:              req.Os,              // 操作系统
		Arch:            req.Arch,            // 体系结构信息
		Engine:          req.Engine,          // 渲染引擎信息
		EngineDetails:   req.EngineDetails,   // 渲染引擎详细信息
		Extra:           req.Extra,           // 其他信息（可选）
		Status:          req.Status,          // 操作状态(0:异常,正常)
		ErrorMsg:        req.ErrorMsg,        // 错误消息
		OperateTime:     req.OperateTime,     // 操作时间
		CostTime:        req.CostTime,        // 消耗时间
	}

	err = r.Service.CreateOperateLog(item)
	if err != nil {
		result.FailWithMsg(c, result.OperateLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteOperateLogByIds 删除操作日志记录
func (r OperateLogController) DeleteOperateLogByIds(c *gin.Context) {

	req := b.DeleteOperateLogReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteOperateLogByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.OperateLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateOperateLog 更新操作日志记录
func (r OperateLogController) UpdateOperateLog(c *gin.Context) {

	req := b.UpdateOperateLogReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateOperateLogDto{
		Id:              req.Id,              // 日志主键
		Title:           req.Title,           // 模块标题
		BusinessType:    req.BusinessType,    // 业务类型（0其它 1新增 2修改 3删除）
		Method:          req.Method,          // 方法名称
		RequestMethod:   req.RequestMethod,   // 请求方式
		OperatorType:    req.OperatorType,    // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     req.OperateName,     // 操作人员
		DeptName:        req.DeptName,        // 部门名称
		OperateUrl:      req.OperateUrl,      // 请求URL
		OperateIp:       req.OperateIp,       // 主机地址
		OperateLocation: req.OperateLocation, // 操作地点
		OperateParam:    req.OperateParam,    // 请求参数
		JsonResult:      req.JsonResult,      // 返回参数
		Platform:        req.Platform,        // 平台信息
		Browser:         req.Browser,         // 浏览器类型
		Version:         req.Version,         // 浏览器版本
		Os:              req.Os,              // 操作系统
		Arch:            req.Arch,            // 体系结构信息
		Engine:          req.Engine,          // 渲染引擎信息
		EngineDetails:   req.EngineDetails,   // 渲染引擎详细信息
		Extra:           req.Extra,           // 其他信息（可选）
		Status:          req.Status,          // 操作状态(0:异常,正常)
		ErrorMsg:        req.ErrorMsg,        // 错误消息
		OperateTime:     req.OperateTime,     // 操作时间
		CostTime:        req.CostTime,        // 消耗时间
	}
	err = r.Service.UpdateOperateLog(item)
	if err != nil {
		result.FailWithMsg(c, result.OperateLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateOperateLogStatus 更新操作日志记录状态
func (r OperateLogController) UpdateOperateLogStatus(c *gin.Context) {

	req := b.UpdateOperateLogStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateOperateLogStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
	}
	err = r.Service.UpdateOperateLogStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.OperateLogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryOperateLogDetail 查询操作日志记录详情
func (r OperateLogController) QueryOperateLogDetail(c *gin.Context) {
	req := b.QueryOperateLogDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryOperateLogDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryOperateLogDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.OperateLogError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryOperateLogList 查询操作日志记录列表
func (r OperateLogController) QueryOperateLogList(c *gin.Context) {
	req := b.QueryOperateLogListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryOperateLogListDto{
		PageNo:          req.PageNo,
		PageSize:        req.PageSize,
		Title:           req.Title,           // 模块标题
		BusinessType:    req.BusinessType,    // 业务类型（0其它 1新增 2修改 3删除）
		Method:          req.Method,          // 方法名称
		RequestMethod:   req.RequestMethod,   // 请求方式
		OperatorType:    req.OperatorType,    // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     req.OperateName,     // 操作人员
		DeptName:        req.DeptName,        // 部门名称
		OperateUrl:      req.OperateUrl,      // 请求URL
		OperateIp:       req.OperateIp,       // 主机地址
		OperateLocation: req.OperateLocation, // 操作地点
		OperateParam:    req.OperateParam,    // 请求参数
		JsonResult:      req.JsonResult,      // 返回参数
		Platform:        req.Platform,        // 平台信息
		Browser:         req.Browser,         // 浏览器类型
		Version:         req.Version,         // 浏览器版本
		Os:              req.Os,              // 操作系统
		Arch:            req.Arch,            // 体系结构信息
		Engine:          req.Engine,          // 渲染引擎信息
		EngineDetails:   req.EngineDetails,   // 渲染引擎详细信息
		Extra:           req.Extra,           // 其他信息（可选）
		Status:          req.Status,          // 操作状态(0:异常,正常)
		ErrorMsg:        req.ErrorMsg,        // 错误消息
		CostTime:        req.CostTime,        // 消耗时间
	}
	list, total := r.Service.QueryOperateLogList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
