package controller

import (
	"github.com/feihua/simple-go/internal/controller/dept"
	"github.com/feihua/simple-go/internal/controller/dict"
	"github.com/feihua/simple-go/internal/controller/log"
	"github.com/feihua/simple-go/internal/controller/menu"
	"github.com/feihua/simple-go/internal/controller/role"
	"github.com/feihua/simple-go/internal/controller/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	dept.NewDeptController,
	dict.NewDictController,
	log.NewLogController,
	menu.NewMenuController,
	role.NewRoleController,
	user.NewUserController,
)
