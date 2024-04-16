package services

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/services/dept"
	"github.com/feihua/simple-go/services/dict"
	"github.com/feihua/simple-go/services/log"
	"github.com/feihua/simple-go/services/menu"
	"github.com/feihua/simple-go/services/role"
	"github.com/feihua/simple-go/services/user"
)

type ServiceImpl struct {
	UserService user.UserService
	RoleService role.RoleService
	MenuService menu.MenuService
	LogService  log.LogService
	DictService dict.DictService
	DeptService dept.DeptService
}

func NewServiceImpl() *ServiceImpl {

	d := dao.Dao
	return &ServiceImpl{
		UserService: user.NewUserServiceImpl(d),
		RoleService: role.NewRoleServiceImpl(d),
		MenuService: menu.NewMenuServiceImpl(d),
		LogService:  log.NewLogServiceImpl(d),
		DictService: dict.NewDictServiceImpl(d),
		DeptService: dept.NewDeptServiceImpl(d),
	}
}

var Service *ServiceImpl

func InitService() {
	Service = NewServiceImpl()

}
