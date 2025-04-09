package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// MenuRouter 菜单信息相关路由
func MenuRouter(r *gin.RouterGroup, b *a.MenuController) {

	r.POST("/dept/addMenu", b.CreateMenu)
	r.POST("/dept/deleteMenuByIds", b.DeleteMenuByIds)
	r.POST("/dept/updateMenu", b.UpdateMenu)
	r.POST("/dept/updateMenuStatus", b.UpdateMenuStatus)
	r.POST("/dept/queryMenuDetail", b.QueryMenuDetail)
	r.POST("/dept/queryMenuList", b.QueryMenuList)

}
