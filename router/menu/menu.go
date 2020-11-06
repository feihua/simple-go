package menu

import (
	"github.com/gin-gonic/gin"
	"simple-go/api/menu"
)

func MenuUrl(r *gin.RouterGroup) {

	r.POST("/menu/add", menu.CreateMenu)
	r.GET("/menu/list", menu.GetMenuList)
	r.POST("/menu/update", menu.UpdateMenu)
	r.POST("/menu/delete", menu.DeleteMenuById)

}
