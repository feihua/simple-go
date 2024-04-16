package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services/dept"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeptController 部门相关操作
/*
Author: LiuFeiHua
Date: 2024/4/15 16:37
*/
type DeptController struct {
}

func NewDeptController() *DeptController {
	return &DeptController{}
}

// CreateDept 添加部门
func (d DeptController) CreateDept(c *gin.Context) {
	req := requests.AddDeptRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service dept.DeptService = &dept.DeptServiceImpl{}

	u := dto.DeptDto{
		DeptName: req.DeptName,
		ParentId: req.ParentId,
		Sort:     req.Sort,
		Remarks:  req.Remarks,
	}

	err = service.CreateDept(u)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}

func (d DeptController) QueryDeptList(c *gin.Context) {
	req := requests.AddDeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptService

	service = &dept.DeptServiceImpl{}

	depts, _ := service.QueryDeptList()
	c.JSON(http.StatusOK, gin.H{"data": depts})
}

func (d DeptController) UpdateDept(c *gin.Context) {

	req := requests.AddDeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptService

	service = &dept.DeptServiceImpl{}
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

func (d DeptController) DeleteDeptByIds(c *gin.Context) {

	req := requests.DeleteDeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dept.DeptService

	service = &dept.DeptServiceImpl{}

	result := service.DeleteDeptByIds(req.Ids)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
