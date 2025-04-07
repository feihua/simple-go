package service

import (
	"github.com/feihua/simple-go/internal/service/system/dept"
	"github.com/feihua/simple-go/internal/service/system/dict"
	"github.com/feihua/simple-go/internal/service/system/log"
	"github.com/feihua/simple-go/internal/service/system/menu"
	"github.com/feihua/simple-go/internal/service/system/role"
	"github.com/feihua/simple-go/internal/service/system/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	dept.NewDeptServiceImpl,
	dict.NewDictServiceImpl,
	log.NewLogServiceImpl,
	menu.NewMenuServiceImpl,
	role.NewRoleServiceImpl,
	user.NewUserServiceImpl,
)
