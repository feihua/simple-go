package menu

import (
	"github.com/feihua/simple-go/api/menu"
	"github.com/gin-gonic/gin"
)

func MenuUrl(r *gin.RouterGroup) {

	r.POST("/menu/add", menu.CreateMenu)
	r.GET("/menu/list", menu.QueryMenuList)
	r.POST("/menu/update", menu.UpdateMenu)
	r.POST("/menu/delete", menu.DeleteMenuById)

}
