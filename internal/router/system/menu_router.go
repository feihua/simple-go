package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// MenuRouter 菜单信息相关路由
func MenuRouter(r *gin.RouterGroup, b *a.MenuController) {

	r.POST("/system/menu/addMenu", b.CreateMenu)
	r.POST("/system/menu/deleteMenu", b.DeleteMenuByIds)
	r.POST("/system/menu/updateMenu", b.UpdateMenu)
	r.POST("/system/menu/updateMenuStatus", b.UpdateMenuStatus)
	r.POST("/system/menu/queryMenuDetail", b.QueryMenuDetail)
	r.POST("/system/menu/queryMenuList", b.QueryMenuList)
	r.GET("/system/menu/queryMenuListSimple", b.QueryMenuListSimple)

}
