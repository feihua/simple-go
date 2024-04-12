package role

import (
	"github.com/feihua/simple-go/api/role"
	"github.com/gin-gonic/gin"
)

func RoleUrl(r *gin.RouterGroup) {

	r.POST("/role/add", role.CreateRole)
	r.GET("/role/list", role.GetRoleList)
	r.POST("/role/update", role.UpdateMenu)
	r.POST("/role/delete", role.DeleteMenuById)
	r.POST("/role/roleMenuIds", role.GeRoleMenuList)
	r.POST("/role/updateRoleMenu", role.UpdateRoleMenu)

}
