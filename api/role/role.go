package role

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go/dto"
	"simple-go/requests"
	"simple-go/services/role"
	"strconv"
)

func CreateRole(c *gin.Context) {

	request := requests.RoleRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}

	u := dto.RoleDto{
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.CreateRole(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetRoleList(c *gin.Context) {
	request := requests.RoleRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service role.RoleContract = &role.RoleService{}

	result, total := service.GetRoleList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func UpdateMenu(c *gin.Context) {

	request := requests.RoleRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}
	u := dto.RoleDto{
		ID:       request.ID,
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.UpdateRole(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteMenuById(c *gin.Context) {

	request := requests.RoleRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}

	result := service.DeleteRoleById(request.ID)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GeRoleMenuList(c *gin.Context) {

	roleId := c.DefaultQuery("roleId", "1")

	var service role.RoleContract = &role.RoleService{}
	menuIds := service.GeRoleMenuList(roleId)
	c.JSON(http.StatusOK, gin.H{"data": menuIds})
}

func UpdateRoleMenu(c *gin.Context) {

	request := requests.RoleMenuRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service role.RoleContract = &role.RoleService{}
	service.UpdateRoleMenu(request)

	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}
