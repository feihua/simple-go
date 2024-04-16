package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/services/log"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (LogController) CreateLog(c *gin.Context) {

	req := requests.SysLogRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogService = &log.LogServiceImpl{}

	u := dto.LogDto{
		UserName:      req.UserName,
		Operation:     req.Operation,
		Method:        req.Method,
		Params:        req.Params,
		OperationTime: req.OperationTime,
		Ip:            req.Ip,
	}
	result := service.CreateLog(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (LogController) QueryLogList(c *gin.Context) {
	req := requests.SysLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service log.LogService = &log.LogServiceImpl{}

	result, total := service.QueryLogList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func (LogController) DeleteLogByIds(c *gin.Context) {

	req := requests.DeleteLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogService = &log.LogServiceImpl{}

	result := service.DeleteLogByIds(req.Ids)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (LogController) CreateLoginLog(c *gin.Context) {

	req := requests.LoginLogRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogService = &log.LogServiceImpl{}

	u := dto.LoginLogDto{
		UserName: req.UserName,
		Ip:       req.Ip,
	}
	result := service.CreateLoginLog(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (LogController) QueryLoginLogList(c *gin.Context) {
	req := requests.LoginLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service log.LogService = &log.LogServiceImpl{}

	result, total := service.QueryLoginLogList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func (LogController) DeleteLoginLogByIds(c *gin.Context) {

	req := requests.DeleteLoginLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogService = &log.LogServiceImpl{}

	result := service.DeleteLoginLogByIds(req.Ids)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
