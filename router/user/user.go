package user

import (
	"github.com/feihua/simple-go/controller/user"
	"github.com/gin-gonic/gin"
)

func UserUrl(r *gin.RouterGroup) {

	r.POST("/user/login", user.Login)
	r.GET("/user/currentUser", user.GetUserInfo)

	r.GET("/hello", user.GetUser)
	r.POST("/user/add", user.CreateUser)
	r.GET("/user/queryUserList", user.QueryUserList)
	r.POST("/user/update", user.UpdateUser)
	r.POST("/user/delete", user.DeleteUserById)
	r.POST("/user/updateUserRole", user.UpdateUserRole)

}
