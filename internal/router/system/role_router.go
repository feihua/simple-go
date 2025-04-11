package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// RoleRouter 角色信息相关路由
func RoleRouter(r *gin.RouterGroup, b *a.RoleController) {

	r.POST("/system/role/addRole", b.CreateRole)
	r.POST("/system/role/deleteRoleByIds", b.DeleteRoleByIds)
	r.POST("/system/role/updateRole", b.UpdateRole)
	r.POST("/system/role/updateRoleStatus", b.UpdateRoleStatus)
	r.POST("/system/role/queryRoleDetail", b.QueryRoleDetail)
	r.POST("/system/role/queryRoleList", b.QueryRoleList)
	r.POST("/system/role/queryRoleMenu", b.QueryRoleMenuList)
	r.POST("/system/role/updateRoleMenu", b.AddRoleMenu)
	r.POST("/system/role/queryAllocatedList", b.QueryAllocatedList)
	r.POST("/system/role/queryUnallocatedList", b.QueryUnallocatedList)
	r.POST("/system/role/cancelAuthUser", b.CancelAuthUser)
	r.POST("/system/role/batchCancelAuthUser", b.BatchCancelAuthUser)
	r.POST("/system/role/batchAuthUser", b.BatchAuthUser)

}
