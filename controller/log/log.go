package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/services/log"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateLog(c *gin.Context) {

	req := requests.SysLogRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

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

func QueryLogList(c *gin.Context) {
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

	var service log.LogContract = &log.LogService{}

	result, total := service.QueryLogList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func DeleteLogById(c *gin.Context) {

	req := requests.SysLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

	result := service.DeleteLogById(req.Id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
