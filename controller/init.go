package controller

import (
	"github.com/feihua/simple-go/controller/dept"
	"github.com/feihua/simple-go/controller/dict"
	"github.com/feihua/simple-go/controller/log"
	"github.com/feihua/simple-go/controller/menu"
	"github.com/feihua/simple-go/controller/role"
	"github.com/feihua/simple-go/controller/user"
	"github.com/feihua/simple-go/services"
)

type ControllerImpl struct {
	UserController *user.UserController
	RoleController *role.RoleController
	MenuController *menu.MenuController
	LogController  *log.LogController
	DictController *dict.DictController
	DeptController *dept.DeptController
}

func NewControllerImpl() *ControllerImpl {
	s := services.Service
	return &ControllerImpl{
		UserController: user.NewUserController(s),
		RoleController: role.NewRoleController(s),
		MenuController: menu.NewMenuController(s),
		LogController:  log.NewLogController(s),
		DictController: dict.NewDictController(s),
		DeptController: dept.NewDeptController(s),
	}
}

var C *ControllerImpl

func InitC() {
	C = NewControllerImpl()
}
