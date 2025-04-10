package system

import (
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/operate_log"
	rq "github.com/feihua/simple-go/internal/vo/system/req"
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

// DeleteOperateLogByIds 删除操作日志记录
func (r OperateLogController) DeleteOperateLogByIds(c *gin.Context) {

	req := rq.DeleteOperateLogReqVo{}
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

// QueryOperateLogDetail 查询操作日志记录详情
func (r OperateLogController) QueryOperateLogDetail(c *gin.Context) {
	req := rq.QueryOperateLogDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryOperateLogDetailDto{
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
	req := rq.QueryOperateLogListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryOperateLogListDto{
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
		Platform:        req.Platform,        // 平台信息
		Browser:         req.Browser,         // 浏览器类型
		Version:         req.Version,         // 浏览器版本
		Os:              req.Os,              // 操作系统
		Status:          req.Status,          // 操作状态(0:异常,正常)

	}
	list, total, err := r.Service.QueryOperateLogList(item)
	if err != nil {
		result.FailWithMsg(c, result.OperateLogError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}
