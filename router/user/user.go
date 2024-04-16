package user

import (
	"github.com/feihua/simple-go/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	controller := user.NewUserController()
	r.POST("/user/login", controller.Login)
	r.GET("/user/info", controller.QueryUserInfo)

	r.POST("/user/addUser", controller.CreateUser)
	r.GET("/user/queryUserList", controller.QueryUserList)
	r.POST("/user/updateUser", controller.UpdateUser)
	r.GET("/user/deleteUserByIds", controller.DeleteUserByIds)
	r.GET("/user/queryUserRoleList", controller.QueryUserRoleList)
	r.POST("/user/updateUserRoleList", controller.UpdateUserRoleList)

}
