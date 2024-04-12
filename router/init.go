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

/*
*
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
