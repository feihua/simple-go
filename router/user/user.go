package user

import (
	"github.com/gin-gonic/gin"
	"simple-go/api/user"
)

func UserUrl(r *gin.RouterGroup) {

	r.POST("/user/login", user.Login)
	r.GET("/user/currentUser", user.GetUserInfo)

	r.GET("/rule/list", user.Rule)
	r.POST("/rule/add", user.AddRule)
	r.POST("/rule/update", user.UpdateRule)
	r.GET("/rule/delete", user.DeleteRule)

	r.GET("/hello", user.GetUser)
	r.POST("/user/add", user.CreateUser)
	r.POST("/user/list", user.GetUserList)
	r.POST("/user/update", user.UpdateUser)
	r.POST("/user/delete", user.DeleteUserById)
	r.POST("/user/updateUserRole", user.UpdateUserRole)

}
