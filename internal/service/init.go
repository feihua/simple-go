package service

import (
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/service/dept"
	"github.com/feihua/simple-go/internal/service/dict"
	"github.com/feihua/simple-go/internal/service/log"
	"github.com/feihua/simple-go/internal/service/menu"
	"github.com/feihua/simple-go/internal/service/role"
	"github.com/feihua/simple-go/internal/service/user"
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
