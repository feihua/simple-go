package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// LoginLogRouter 系统访问记录相关路由
func LoginLogRouter(r *gin.RouterGroup, b *a.LoginLogController) {

	r.POST("/system/loginLog/addLoginLog", b.CreateLoginLog)
	r.POST("/system/loginLog/deleteLoginLogByIds", b.DeleteLoginLogByIds)
	r.POST("/system/loginLog/updateLoginLog", b.UpdateLoginLog)
	r.POST("/system/loginLog/updateLoginLogStatus", b.UpdateLoginLogStatus)
	r.POST("/system/loginLog/queryLoginLogDetail", b.QueryLoginLogDetail)
	r.POST("/system/loginLog/queryLoginLogList", b.QueryLoginLogList)

}
