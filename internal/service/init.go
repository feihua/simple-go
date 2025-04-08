package service

import (
	"github.com/feihua/simple-go/internal/service/system/dept"
	"github.com/feihua/simple-go/internal/service/system/dict_data"
	"github.com/feihua/simple-go/internal/service/system/dict_type"
	"github.com/feihua/simple-go/internal/service/system/login_log"
	"github.com/feihua/simple-go/internal/service/system/menu"
	"github.com/feihua/simple-go/internal/service/system/notice"
	"github.com/feihua/simple-go/internal/service/system/operate_log"
	"github.com/feihua/simple-go/internal/service/system/post"
	"github.com/feihua/simple-go/internal/service/system/role"
	"github.com/feihua/simple-go/internal/service/system/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	dept.NewDeptServiceImpl,
	dict_data.NewDictDataServiceImpl,
	dict_type.NewDictTypeServiceImpl,
	login_log.NewLoginLogServiceImpl,
	menu.NewMenuServiceImpl,
	notice.NewNoticeServiceImpl,
	operate_log.NewOperateLogServiceImpl,
	post.NewPostServiceImpl,
	role.NewRoleServiceImpl,
	user.NewUserServiceImpl,
)
