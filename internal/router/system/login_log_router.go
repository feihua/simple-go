package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// LoginLogRouter 系统访问记录相关路由
func LoginLogRouter(r *gin.RouterGroup, b *a.LoginLogController) {

	r.POST("/dept/addLoginLog", b.CreateLoginLog)
	r.POST("/dept/deleteLoginLogByIds", b.DeleteLoginLogByIds)
	r.POST("/dept/updateLoginLog", b.UpdateLoginLog)
	r.POST("/dept/updateLoginLogStatus", b.UpdateLoginLogStatus)
	r.POST("/dept/queryLoginLogDetail", b.QueryLoginLogDetail)
	r.POST("/dept/queryLoginLogList", b.QueryLoginLogList)

}
