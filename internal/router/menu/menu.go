package menu

import (
	"github.com/feihua/simple-go/internal/controller"
	"github.com/gin-gonic/gin"
)

func MenuRouter(r *gin.RouterGroup) {

	menuController := controller.C.MenuController
	r.POST("/menu/addMenu", menuController.CreateMenu)
	r.GET("/menu/queryMenuList", menuController.QueryMenuList)
	r.POST("/menu/updateMenu", menuController.UpdateMenu)
	r.POST("/menu/deleteMenuById", menuController.DeleteMenuById)

}
