package log

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go/dto"
	"simple-go/requests"
	"simple-go/services/log"
	"strconv"
)

func CreateLoginLog(c *gin.Context) {

	request := requests.LoginLogRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

	u := dto.LoginLogDto{
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.CreateLoginLog(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetLoginLogList(c *gin.Context) {
	request := requests.LoginLogRequest{}
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

	result, total := service.GetLoginLogList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func UpdateLoginLog(c *gin.Context) {

	request := requests.LoginLogRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}
	u := dto.LoginLogDto{
		ID:       request.ID,
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.UpdateLoginLog(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteLoginLogById(c *gin.Context) {

	request := requests.LoginLogRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service log.LogContract = &log.LogService{}

	result := service.DeleteLoginLogById(request.ID)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
