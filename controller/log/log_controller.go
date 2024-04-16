package log

import (
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services/log"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"strconv"
)

// LogController 日志相关
/*
Author: LiuFeiHua
Date: 2024/4/15 18:02
*/
type LogController struct {
}

func NewLogController() *LogController {
	return &LogController{}
}

// QueryLogList 查询操作日志
func (LogController) QueryLogList(c *gin.Context) {
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

	var service log.LogService = &log.LogServiceImpl{}

	list, total := service.QueryLogList(pageNum, size)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

// DeleteLogByIds 删除操作日志
func (LogController) DeleteLogByIds(c *gin.Context) {

	req := requests.DeleteLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service log.LogService = &log.LogServiceImpl{}

	err = service.DeleteLogByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.LogError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryLoginLogList 查询登录日志
func (LogController) QueryLoginLogList(c *gin.Context) {
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

	var service log.LogService = &log.LogServiceImpl{}

	list, total := service.QueryLoginLogList(pageNum, size)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

// DeleteLoginLogByIds 删除登录日志
func (LogController) DeleteLoginLogByIds(c *gin.Context) {

	req := requests.DeleteLoginLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service log.LogService = &log.LogServiceImpl{}

	err = service.DeleteLoginLogByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.LogError, err.Error())
	} else {
		result.Ok(c)
	}
}
