package menu

import (
	"github.com/feihua/simple-go/controller/menu"
	"github.com/gin-gonic/gin"
)

func MenuRouter(r *gin.RouterGroup) {

	controller := menu.NewMenuController()
	r.POST("/menu/addMenu", controller.CreateMenu)
	r.GET("/menu/queryMenuList", controller.QueryMenuList)
	r.POST("/menu/updateMenu", controller.UpdateMenu)
	r.POST("/menu/deleteMenuById", controller.DeleteMenuById)

}
