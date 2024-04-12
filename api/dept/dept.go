package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/requests"
	"github.com/feihua/simple-go/services/dept"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDept(c *gin.Context) {

	req := requests.DeptRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract = &dept.DeptService{}

	u := dto.DeptDto{
		DeptName: req.DeptName,
		ParentId: req.ParentId,
		Sort:     req.Sort,
		Remarks:  req.Remarks,
	}
	result := service.CreateDept(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func QueryDeptList(c *gin.Context) {
	req := requests.DeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract

	service = &dept.DeptService{}

	depts, _ := service.QueryDeptList()
	c.JSON(http.StatusOK, gin.H{"data": depts})
}

func UpdateDept(c *gin.Context) {

	req := requests.DeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract

	service = &dept.DeptService{}
	u := dto.DeptDto{
		Id:       req.Id,
		DeptName: req.DeptName,
		ParentId: req.ParentId,
		Sort:     req.Sort,
		Remarks:  req.Remarks,
	}
	result := service.UpdateDept(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteDeptById(c *gin.Context) {

	req := requests.DeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptContract

	service = &dept.DeptService{}

	result := service.DeleteDeptById(req.Id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
