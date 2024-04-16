package role

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services/role"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RoleController 角色相关
/*
Author: LiuFeiHua
Date: 2024/4/15 16:55
*/
type RoleController struct {
}

func NewRoleController() *RoleController {
	return &RoleController{}
}
func (RoleController) CreateRole(c *gin.Context) {

	req := requests.RoleRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service role.RoleService = &role.RoleServiceImpl{}

	u := dto.RoleDto{
		RoleName: req.RoleName,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}
	result := service.CreateRole(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (RoleController) QueryRoleList(c *gin.Context) {
	req := requests.RoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service role.RoleService = &role.RoleServiceImpl{}

	result, total := service.QueryRoleList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func (RoleController) UpdateMenu(c *gin.Context) {

	req := requests.RoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service role.RoleService = &role.RoleServiceImpl{}
	u := dto.RoleDto{
		Id:       req.Id,
		RoleName: req.RoleName,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}
	result := service.UpdateRole(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (RoleController) DeleteRoleByIds(c *gin.Context) {

	req := requests.DeleteRoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service role.RoleService = &role.RoleServiceImpl{}

	result := service.DeleteRoleByIds(req.Ids)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (RoleController) QueryRoleMenuList(c *gin.Context) {

	roleId := c.DefaultQuery("roleId", "1")

	var service role.RoleService = &role.RoleServiceImpl{}
	menuIds := service.QueryRoleMenuList(roleId)
	c.JSON(http.StatusOK, gin.H{"data": menuIds})
}

func (RoleController) UpdateRoleMenuList(c *gin.Context) {

	req := requests.RoleMenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service role.RoleService = &role.RoleServiceImpl{}
	service.UpdateRoleMenu(req)

	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}
