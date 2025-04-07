package service

import (
	"github.com/feihua/simple-go/internal/service/dept"
	"github.com/feihua/simple-go/internal/service/dict"
	"github.com/feihua/simple-go/internal/service/log"
	"github.com/feihua/simple-go/internal/service/menu"
	"github.com/feihua/simple-go/internal/service/role"
	"github.com/feihua/simple-go/internal/service/user"
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
