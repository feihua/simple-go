package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/dept"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// DeptController 部门相关
type DeptController struct {
	Service dept.DeptService
}

func NewDeptController(Service dept.DeptService) *DeptController {
	return &DeptController{Service: Service}
}

// CreateDept 添加部门
func (r DeptController) CreateDept(c *gin.Context) {

	req := b.AddDeptReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.AddDeptDto{
		ParentId:  req.ParentId,  // 父部门id
		Ancestors: req.Ancestors, // 祖级列表
		DeptName:  req.DeptName,  // 部门名称
		Sort:      req.Sort,      // 显示顺序
		Leader:    req.Leader,    // 负责人
		Phone:     req.Phone,     // 联系电话
		Email:     req.Email,     // 邮箱
		Status:    req.Status,    // 部门状态（0：停用，1:正常）
		CreateBy:  c.MustGet("userName").(string),
	}

	err = r.Service.CreateDept(item)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteDeptByIds 删除部门
func (r DeptController) DeleteDeptByIds(c *gin.Context) {

	req := b.DeleteDeptReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteDeptByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateDept 更新部门
func (r DeptController) UpdateDept(c *gin.Context) {

	req := b.UpdateDeptReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateDeptDto{
		Id:        req.Id,        // 部门id
		ParentId:  req.ParentId,  // 父部门id
		Ancestors: req.Ancestors, // 祖级列表
		DeptName:  req.DeptName,  // 部门名称
		Sort:      req.Sort,      // 显示顺序
		Leader:    req.Leader,    // 负责人
		Phone:     req.Phone,     // 联系电话
		Email:     req.Email,     // 邮箱
		Status:    req.Status,    // 部门状态（0：停用，1:正常）
		CreateBy:  c.MustGet("userName").(string),
	}
	err = r.Service.UpdateDept(item)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateDeptStatus 更新部门状态
func (r DeptController) UpdateDeptStatus(c *gin.Context) {

	req := b.UpdateDeptStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateDeptStatusDto{
		Ids:      req.Ids,
		Status:   req.Status,
		UpdateBy: c.MustGet("userName").(string),
	}
	err = r.Service.UpdateDeptStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryDeptDetail 查询部门详情
func (r DeptController) QueryDeptDetail(c *gin.Context) {
	req := b.QueryDeptDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryDeptDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryDeptDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryDeptList 查询部门列表
func (r DeptController) QueryDeptList(c *gin.Context) {
	req := b.QueryDeptListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryDeptListDto{}
	list, err := r.Service.QueryDeptList(item)
	if err != nil {
		result.FailWithMsg(c, result.DeptError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true})
	}

}
