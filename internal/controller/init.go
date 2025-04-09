package controller

import (
	"github.com/feihua/simple-go/internal/controller/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	system.NewDeptController,
	system.NewDictTypeController,
	system.NewDictDataController,
	system.NewLoginLogController,
	system.NewMenuController,
	system.NewNoticeController,
	system.NewOperateLogController,
	system.NewPostController,
	system.NewRoleController,
	system.NewUserController,
)
