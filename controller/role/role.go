package role

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/services/role"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateRole(c *gin.Context) {

	req := requests.RoleRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}

	u := dto.RoleDto{
		RoleName: req.RoleName,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}
	result := service.CreateRole(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func QueryRoleList(c *gin.Context) {
	req := requests.RoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service role.RoleContract = &role.RoleService{}

	result, total := service.QueryRoleList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func UpdateMenu(c *gin.Context) {

	req := requests.RoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}
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

func DeleteMenuById(c *gin.Context) {

	req := requests.RoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}

	result := service.DeleteRoleById(req.Id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func QueryRoleMenuList(c *gin.Context) {

	roleId := c.DefaultQuery("roleId", "1")

	var service role.RoleContract = &role.RoleService{}
	menuIds := service.QueryRoleMenuList(roleId)
	c.JSON(http.StatusOK, gin.H{"data": menuIds})
}

func UpdateRoleMenu(c *gin.Context) {

	req := requests.RoleMenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}
	service.UpdateRoleMenu(req)

	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}
