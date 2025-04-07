package log

import (
	"github.com/feihua/simple-go/internal/services"
	"github.com/feihua/simple-go/internal/vo/requests"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
	"strconv"
)

// LogController 日志相关
/*
Author: LiuFeiHua
Date: 2024/4/15 18:02
*/
type LogController struct {
	Service *services.ServiceImpl
}

func NewLogController(Service *services.ServiceImpl) *LogController {
	return &LogController{Service: Service}
}

// QueryLogList 查询操作日志
func (l LogController) QueryLogList(c *gin.Context) {
	req := requests.SysLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	list, total := l.Service.LogService.QueryLogList(pageNum, size)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

// DeleteLogByIds 删除操作日志
func (l LogController) DeleteLogByIds(c *gin.Context) {

	req := requests.DeleteLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = l.Service.LogService.DeleteLogByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.LogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryLoginLogList 查询登录日志
func (l LogController) QueryLoginLogList(c *gin.Context) {
	req := requests.LoginLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	list, total := l.Service.LogService.QueryLoginLogList(pageNum, size)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

// DeleteLoginLogByIds 删除登录日志
func (l LogController) DeleteLoginLogByIds(c *gin.Context) {

	req := requests.DeleteLoginLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = l.Service.LogService.DeleteLoginLogByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.LogError, err.Error())
	} else {
		result.Ok(c)
	}
}
