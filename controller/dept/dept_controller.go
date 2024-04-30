package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
)

// DeptController 部门相关操作
/*
Author: LiuFeiHua
Date: 2024/4/15 16:37
*/
type DeptController struct {
	Service *services.ServiceImpl
}

func NewDeptController(Service *services.ServiceImpl) *DeptController {
	return &DeptController{Service: Service}
}

// CreateDept 添加部门
func (d DeptController) CreateDept(c *gin.Context) {
	req := requests.DeptRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	deptDto := dto.DeptDto{
		DeptName: req.DeptName,
		ParentId: req.ParentId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}

	err = d.Service.DeptService.CreateDept(deptDto)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryDeptList 查询部门列表
func (d DeptController) QueryDeptList(c *gin.Context) {
	req := requests.DeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	deptList, err := d.Service.DeptService.QueryDeptList()
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.OkWithData(c, deptList)
	}
}

// UpdateDept 修改部门
func (d DeptController) UpdateDept(c *gin.Context) {

	req := requests.DeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	deptDto := dto.DeptDto{
		Id:       req.Id,
		DeptName: req.DeptName,
		ParentId: req.ParentId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}

	err = d.Service.DeptService.UpdateDept(deptDto)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteDeptByIds 删除部门
func (d DeptController) DeleteDeptByIds(c *gin.Context) {

	req := requests.DeleteDeptRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = d.Service.DeptService.DeleteDeptByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}
