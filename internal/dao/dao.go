package dao

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewDeptDao,
	NewDictDao,
	NewLogDao,
	NewMenuDao,
	NewRoleMenuDao,
	NewRoleDao,
	NewUserDao,
	NewUserRoleDao,
)
