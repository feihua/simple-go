package controller

import (
	"github.com/feihua/simple-go/internal/controller/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	system.NewDeptController,
	system.NewDictController,
	system.NewLogController,
	system.NewMenuController,
	system.NewRoleController,
	system.NewUserController,
)
