package dao

import (
	"github.com/feihua/simple-go/internal/dao/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	system.NewDeptDao,
	system.NewDictTypeDao,
	system.NewDictDataDao,
	system.NewLoginLogDao,
	system.NewMenuDao,
	system.NewNoticeDao,
	system.NewOperateLogDao,
	system.NewPostDao,
	system.NewRoleDao,
	system.NewRoleDeptDao,
	system.NewRoleMenuDao,
	system.NewUserDao,
	system.NewUserPostDao,
	system.NewUserRoleDao,
)
