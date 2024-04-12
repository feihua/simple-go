package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/requests"
	"github.com/feihua/simple-go/services/dept"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDept(c *gin.Context) {

	request := requests.DeptRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract = &dept.DeptService{}

	u := dto.DeptDto{
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.CreateDept(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetDeptList(c *gin.Context) {
	request := requests.DeptRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract

	service = &dept.DeptService{}

	result := service.GetDeptList()
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateDept(c *gin.Context) {

	request := requests.DeptRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract

	service = &dept.DeptService{}
	u := dto.DeptDto{
		ID:       request.ID,
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.UpdateDept(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteDeptById(c *gin.Context) {

	request := requests.DeptRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract

	service = &dept.DeptService{}

	result := service.DeleteDeptById(request.ID)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
