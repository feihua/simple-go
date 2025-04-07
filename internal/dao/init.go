package dao

import (
	"github.com/feihua/simple-go/internal/dao/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	system.NewDeptDao,
	system.NewDictDao,
	system.NewLogDao,
	system.NewMenuDao,
	system.NewRoleMenuDao,
	system.NewRoleDao,
	system.NewUserDao,
	system.NewUserRoleDao,
)
