package user

import (
	"github.com/feihua/simple-go/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	userController := controller.C.UserController
	r.POST("/user/login", userController.Login)
	r.GET("/user/queryUserMenuList", userController.QueryUserMenuList)

	r.POST("/user/addUser", userController.CreateUser)
	r.POST("/user/queryUserList", userController.QueryUserList)
	r.POST("/user/updateUser", userController.UpdateUser)
	r.POST("/user/deleteUserByIds", userController.DeleteUserByIds)
	r.GET("/user/queryUserRoleList", userController.QueryUserRoleList)
	r.POST("/user/updateUserRoleList", userController.UpdateUserRoleList)

}
