package log

import (
	"github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup, logController *system.LogController) {

	r.POST("/sysLog/queryOperateLogList", logController.QueryLogList)
	r.POST("/sysLog/deleteOperateLogByIds", logController.DeleteLogByIds)

	r.POST("/loginLog/queryLoginLogList", logController.QueryLoginLogList)
	r.POST("/loginLog/deleteLoginLogByIds", logController.DeleteLoginLogByIds)
}
