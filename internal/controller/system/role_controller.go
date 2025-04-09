package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/role"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// RoleController 角色信息相关
type RoleController struct {
	Service role.RoleService
}

func NewRoleController(Service role.RoleService) *RoleController {
	return &RoleController{Service: Service}
}

// CreateRole 添加角色信息
func (r RoleController) CreateRole(c *gin.Context) {

	req := b.AddRoleReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.AddRoleDto{
		RoleName:  req.RoleName,  // 名称
		RoleKey:   req.RoleKey,   // 角色权限字符串
		DataScope: req.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,    // 状态(1:正常，0:禁用)
		Remark:    req.Remark,    // 备注

	}

	err = r.Service.CreateRole(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteRoleByIds 删除角色信息
func (r RoleController) DeleteRoleByIds(c *gin.Context) {

	req := b.DeleteRoleReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteRoleByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateRole 更新角色信息
func (r RoleController) UpdateRole(c *gin.Context) {

	req := b.UpdateRoleReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateRoleDto{
		Id:        req.Id,        // 主键
		RoleName:  req.RoleName,  // 名称
		RoleKey:   req.RoleKey,   // 角色权限字符串
		DataScope: req.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,    // 状态(1:正常，0:禁用)
		Remark:    req.Remark,    // 备注

	}
	err = r.Service.UpdateRole(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateRoleStatus 更新角色信息状态
func (r RoleController) UpdateRoleStatus(c *gin.Context) {

	req := b.UpdateRoleStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateRoleStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
	}
	err = r.Service.UpdateRoleStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryRoleDetail 查询角色信息详情
func (r RoleController) QueryRoleDetail(c *gin.Context) {
	req := b.QueryRoleDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryRoleDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryRoleDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryRoleList 查询角色信息列表
func (r RoleController) QueryRoleList(c *gin.Context) {
	req := b.QueryRoleListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryRoleListDto{
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
		RoleName:  req.RoleName,  // 名称
		RoleKey:   req.RoleKey,   // 角色权限字符串
		DataScope: req.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,    // 状态(1:正常，0:禁用)
	}
	list, total := r.Service.QueryRoleList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
