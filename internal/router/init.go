package router

import (
	de1 "github.com/feihua/simple-go/internal/controller/dept"
	d1 "github.com/feihua/simple-go/internal/controller/dict"
	l1 "github.com/feihua/simple-go/internal/controller/log"
	m1 "github.com/feihua/simple-go/internal/controller/menu"
	r1 "github.com/feihua/simple-go/internal/controller/role"
	u1 "github.com/feihua/simple-go/internal/controller/user"
	"github.com/feihua/simple-go/internal/middleware"
	"github.com/feihua/simple-go/internal/router/dept"
	"github.com/feihua/simple-go/internal/router/dict"
	"github.com/feihua/simple-go/internal/router/log"
	"github.com/feihua/simple-go/internal/router/menu"
	"github.com/feihua/simple-go/internal/router/role"
	"github.com/feihua/simple-go/internal/router/user"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(uController *u1.UserController, rController *r1.RoleController, mController *m1.MenuController, lController *l1.LogController, dController *d1.DictController, deController *de1.DeptController) *gin.Engine {
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
