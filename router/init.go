package router

import (
	"github.com/gin-gonic/gin"
	"simple-go/router/dept"
	"simple-go/router/dict"
	"simple-go/router/log"
	"simple-go/router/menu"
	"simple-go/router/role"
	"simple-go/router/user"
)

/**
初始化路由
*/
func Init(r *gin.RouterGroup) {

	//用户
	user.UserUrl(r)
	//角色
	role.RoleUrl(r)
	//菜单
	menu.MenuUrl(r)
	//日志
	log.LogUrl(r)
	//字典
	dict.DictUrl(r)
	//部门
	dept.DeptUrl(r)
}
