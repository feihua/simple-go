package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户信息相关路由
func UserRouter(r *gin.RouterGroup, b *a.UserController) {

	r.POST("/system/user/addUser", b.CreateUser)
	r.POST("/system/user/deleteUser", b.DeleteUserByIds)
	r.POST("/system/user/updateUser", b.UpdateUser)
	r.POST("/system/user/updateUserStatus", b.UpdateUserStatus)
	r.POST("/system/user/queryUserDetail", b.QueryUserDetail)
	r.POST("/system/user/queryUserList", b.QueryUserList)
	r.POST("/system/user/login", b.Login)
	r.GET("/system/user/queryUserMenu", b.QueryUserMenuList)
	r.POST("/system/user/queryUserRole", b.QueryUserRoleList)
	r.POST("/system/user/updateUserRole", b.UpdateUserRoleList)

}
