package router

import (
	u1 "github.com/feihua/simple-go/internal/controller/system"
	"github.com/feihua/simple-go/internal/middleware"
	"github.com/feihua/simple-go/internal/router/system/dept"
	"github.com/feihua/simple-go/internal/router/system/dict"
	"github.com/feihua/simple-go/internal/router/system/log"
	"github.com/feihua/simple-go/internal/router/system/menu"
	"github.com/feihua/simple-go/internal/router/system/role"
	"github.com/feihua/simple-go/internal/router/system/user"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(uController *u1.UserController, rController *u1.RoleController, mController *u1.MenuController, lController *u1.LogController, dController *u1.DictController, deController *u1.DeptController) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JwtMiddleware())

	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	routerGroup := r.Group("/api")
	// 用户
	user.UserRouter(routerGroup, uController)
	// 角色
	role.RoleRouter(routerGroup, rController)
	// 菜单
	menu.MenuRouter(routerGroup, mController)
	// 日志
	log.LogRouter(routerGroup, lController)
	// 字典
	dict.DictRouter(routerGroup, dController)
	// 部门
	dept.DeptRouter(routerGroup, deController)

	return r
}
