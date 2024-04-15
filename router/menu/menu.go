package menu

import (
	"github.com/feihua/simple-go/controller/menu"
	"github.com/gin-gonic/gin"
)

func MenuUrl(r *gin.RouterGroup) {

	controller := menu.NewMenuController()
	r.POST("/menu/addMenu", controller.CreateMenu)
	r.GET("/menu/queryMenuList", controller.QueryMenuList)
	r.POST("/menu/updateMenu", controller.UpdateMenu)
	r.POST("/menu/deleteMenuByIds", controller.DeleteMenuByIds)

}
