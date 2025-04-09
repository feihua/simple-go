package router

import (
	u1 "github.com/feihua/simple-go/internal/controller/system"
	"github.com/feihua/simple-go/internal/middleware"
	"github.com/feihua/simple-go/internal/router/system"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(user *u1.UserController, role *u1.RoleController, menu *u1.MenuController, loginLog *u1.LoginLogController,
	operateLog *u1.OperateLogController, dictData *u1.DictDataController, dictType *u1.DictTypeController, notice *u1.NoticeController,
	dept *u1.DeptController, post *u1.PostController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.JwtMiddleware())

	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	routerGroup := r.Group("/api")
	system.UserRouter(routerGroup, user)
	system.RoleRouter(routerGroup, role)
	system.MenuRouter(routerGroup, menu)
	system.LoginLogRouter(routerGroup, loginLog)
	system.OperateLogRouter(routerGroup, operateLog)
	system.DictDataRouter(routerGroup, dictData)
	system.DictTypeRouter(routerGroup, dictType)
	system.NoticeRouter(routerGroup, notice)
	system.DeptRouter(routerGroup, dept)
	system.PostRouter(routerGroup, post)

	return r
}
