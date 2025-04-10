package system

import (
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/login_log"
	rq "github.com/feihua/simple-go/internal/vo/system/req"
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

// DeleteLoginLogByIds 删除系统访问记录
func (r LoginLogController) DeleteLoginLogByIds(c *gin.Context) {

	req := rq.DeleteLoginLogReqVo{}
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

// QueryLoginLogDetail 查询系统访问记录详情
func (r LoginLogController) QueryLoginLogDetail(c *gin.Context) {
	req := rq.QueryLoginLogDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryLoginLogDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryLoginLogDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.LoginLogError, err.Error())
	} else {
		result.OkWithData(c, data)
	}
}

// QueryLoginLogList 查询系统访问记录列表
func (r LoginLogController) QueryLoginLogList(c *gin.Context) {
	req := rq.QueryLoginLogListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryLoginLogListDto{
		PageNo:        req.PageNo,
		PageSize:      req.PageSize,
		LoginName:     req.LoginName,     // 登录账号
		Ipaddr:        req.Ipaddr,        // 登录IP地址
		LoginLocation: req.LoginLocation, // 登录地点
		Platform:      req.Platform,      // 平台信息
		Browser:       req.Browser,       // 浏览器类型
		Version:       req.Version,       // 浏览器版本
		Os:            req.Os,            // 操作系统
		Status:        req.Status,        // 登录状态(0:失败,1:成功)
	}
	list, total, err := r.Service.QueryLoginLogList(item)
	if err != nil {
		result.FailWithMsg(c, result.LoginLogError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}
