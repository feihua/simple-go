package role

import (
	"github.com/feihua/simple-go/controller/role"
	"github.com/gin-gonic/gin"
)

func RoleUrl(r *gin.RouterGroup) {

	controller := role.NewRoleController()
	r.POST("/role/addRole", controller.CreateRole)
	r.GET("/role/queryRoleList", controller.QueryRoleList)
	r.POST("/role/updateRole", controller.UpdateMenu)
	r.POST("/role/deleteRoleByIds", controller.DeleteRoleByIds)
	r.POST("/role/queryRoleMenuList", controller.QueryRoleMenuList)
	r.POST("/role/updateRoleMenuList", controller.UpdateRoleMenuList)

}
