package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户信息相关路由
func UserRouter(r *gin.RouterGroup, b *a.UserController) {

	r.POST("/dept/addUser", b.CreateUser)
	r.POST("/dept/deleteUserByIds", b.DeleteUserByIds)
	r.POST("/dept/updateUser", b.UpdateUser)
	r.POST("/dept/updateUserStatus", b.UpdateUserStatus)
	r.POST("/dept/queryUserDetail", b.QueryUserDetail)
	r.POST("/dept/queryUserList", b.QueryUserList)

}
