package router

import (
	"github.com/feihua/simple-go/router/dept"
	"github.com/feihua/simple-go/router/dict"
	"github.com/feihua/simple-go/router/log"
	"github.com/feihua/simple-go/router/menu"
	"github.com/feihua/simple-go/router/role"
	"github.com/feihua/simple-go/router/user"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(r *gin.RouterGroup) {

	//用户
	user.UserRouter(r)
	//角色
	role.RoleRouter(r)
	//菜单
	menu.MenuRouter(r)
	//日志
	log.LogRouter(r)
	//字典
	dict.DictRouter(r)
	//部门
	dept.DeptRouter(r)
}
