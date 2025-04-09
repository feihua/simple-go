package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// RoleRouter 角色信息相关路由
func RoleRouter(r *gin.RouterGroup, b *a.RoleController) {

	r.POST("/dept/addRole", b.CreateRole)
	r.POST("/dept/deleteRoleByIds", b.DeleteRoleByIds)
	r.POST("/dept/updateRole", b.UpdateRole)
	r.POST("/dept/updateRoleStatus", b.UpdateRoleStatus)
	r.POST("/dept/queryRoleDetail", b.QueryRoleDetail)
	r.POST("/dept/queryRoleList", b.QueryRoleList)

}
