package role

import (
	"github.com/feihua/simple-go/internal/controller/role"
	"github.com/gin-gonic/gin"
)

func RoleRouter(r *gin.RouterGroup, roleController *role.RoleController) {

	r.POST("/role/addRole", roleController.CreateRole)
	r.POST("/role/queryRoleList", roleController.QueryRoleList)
	r.POST("/role/updateRole", roleController.UpdateRole)
	r.POST("/role/deleteRoleByIds", roleController.DeleteRoleByIds)
	r.GET("/role/queryRoleMenuList", roleController.QueryRoleMenuList)
	r.POST("/role/updateRoleMenuList", roleController.UpdateRoleMenuList)

}
