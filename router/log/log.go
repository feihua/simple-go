package log

import (
	"github.com/feihua/simple-go/controller"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {

	logController := controller.C.LogController
	r.GET("/syslog/queryLogList", logController.QueryLogList)
	r.POST("/syslog/deleteLogByIds", logController.DeleteLogByIds)

	r.GET("/loginLog/queryLoginLogList", logController.QueryLoginLogList)
	r.POST("/loginLog/deleteLoginLogByIds", logController.DeleteLoginLogByIds)
}
