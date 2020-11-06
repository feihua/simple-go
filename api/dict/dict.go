package dict

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go/dto"
	"simple-go/requests"
	"simple-go/services/dict"
	"strconv"
)

func CreateDict(c *gin.Context) {

	request := requests.DictRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dict.DictContract = &dict.DictService{}

	u := dto.DictDto{
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.CreateDict(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetDictList(c *gin.Context) {
	request := requests.DictRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service dict.DictContract = &dict.DictService{}

	result, total := service.GetDictList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func UpdateDict(c *gin.Context) {

	request := requests.DictRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dict.DictContract = &dict.DictService{}

	u := dto.DictDto{
		ID:       request.ID,
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.UpdateDict(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteDictById(c *gin.Context) {

	request := requests.DictRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dict.DictContract = &dict.DictService{}

	result := service.DeleteDictById(request.ID)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
