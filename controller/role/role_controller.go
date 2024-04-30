package role

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"strconv"
)

// RoleController 角色相关
/*
Author: LiuFeiHua
Date: 2024/4/15 16:55
*/
type RoleController struct {
	Service *services.ServiceImpl
}

func NewRoleController(Service *services.ServiceImpl) *RoleController {
	return &RoleController{Service: Service}
}

// CreateRole 创建角色
func (r RoleController) CreateRole(c *gin.Context) {

	req := requests.RoleRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	roleDto := dto.RoleDto{
		RoleName: req.RoleName,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}

	err = r.Service.RoleService.CreateRole(roleDto)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryRoleList 查询角色列表
func (r RoleController) QueryRoleList(c *gin.Context) {
	req := requests.QueryRoleListRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	roleListDto := dto.QueryRoleListDto{
		RoleName: req.RoleName,
		StatusId: req.StatusId,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	}
	list, total := r.Service.RoleService.QueryRoleList(roleListDto)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}

// UpdateRole 更新角色
func (r RoleController) UpdateRole(c *gin.Context) {

	req := requests.RoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	roleDto := dto.RoleDto{
		Id:       req.Id,
		RoleName: req.RoleName,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}
	err = r.Service.RoleService.UpdateRole(roleDto)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteRoleByIds 删除角色
func (r RoleController) DeleteRoleByIds(c *gin.Context) {

	req := requests.DeleteRoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = r.Service.RoleService.DeleteRoleByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryRoleMenuList 根据角色Id查询角色菜单
func (r RoleController) QueryRoleMenuList(c *gin.Context) {

	roleId := c.DefaultQuery("roleId", "10")
	id, err := strconv.ParseInt(roleId, 10, 64)

	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	resp, err := r.Service.RoleService.QueryRoleMenuList(id)

	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.OkWithData(c, resp)
	}
}

// UpdateRoleMenuList 更新角色菜单
func (r RoleController) UpdateRoleMenuList(c *gin.Context) {

	req := requests.RoleMenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	menuDtoRequest := dto.RoleMenuDtoRequest{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	}
	err = r.Service.RoleService.UpdateRoleMenuList(menuDtoRequest)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}
