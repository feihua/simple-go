package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/services/log"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateLoginLog(c *gin.Context) {

	req := requests.LoginLogRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

	u := dto.LoginLogDto{
		UserName: req.UserName,
		Ip:       req.Ip,
	}
	result := service.CreateLoginLog(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func QueryLoginLogList(c *gin.Context) {
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

	var service log.LogContract = &log.LogService{}

	result, total := service.QueryLoginLogList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func DeleteLoginLogById(c *gin.Context) {

	req := requests.LoginLogRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

	result := service.DeleteLoginLogById(req.Id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
