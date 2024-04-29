package user

import (
	"github.com/feihua/simple-go/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	userController := controller.C.UserController
	r.POST("/user/login", userController.Login)
	r.GET("/user/queryUserMenu", userController.QueryUserMenu)

	r.POST("/user/addUser", userController.CreateUser)
	r.POST("/user/queryUserList", userController.QueryUserList)
	r.POST("/user/updateUser", userController.UpdateUser)
	r.GET("/user/deleteUserByIds", userController.DeleteUserByIds)
	r.GET("/user/queryUserRoleList", userController.QueryUserRoleList)
	r.POST("/user/updateUserRoleList", userController.UpdateUserRoleList)

}
