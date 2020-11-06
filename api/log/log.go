package log

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go/dto"
	"simple-go/requests"
	"simple-go/services/log"
	"strconv"
)

func CreateLog(c *gin.Context) {

	request := requests.SysLogRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

	u := dto.LogDto{
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.CreateLog(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetLogList(c *gin.Context) {
	request := requests.SysLogRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service log.LogContract = &log.LogService{}

	result, total := service.GetLogList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func UpdateLog(c *gin.Context) {

	request := requests.SysLogRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}
	u := dto.LogDto{
		ID:       request.ID,
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.UpdateLog(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteLogById(c *gin.Context) {

	request := requests.SysLogRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

	result := service.DeleteLogById(request.ID)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
