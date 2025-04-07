package services

import (
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/services/dept"
	"github.com/feihua/simple-go/internal/services/dict"
	"github.com/feihua/simple-go/internal/services/log"
	"github.com/feihua/simple-go/internal/services/menu"
	"github.com/feihua/simple-go/internal/services/role"
	"github.com/feihua/simple-go/internal/services/user"
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
