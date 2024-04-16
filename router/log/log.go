package log

import (
	"github.com/feihua/simple-go/controller/log"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {

	controller := log.NewLogController()
	r.GET("/syslog/queryLogList", controller.QueryLogList)
	r.POST("/syslog/deleteLogByIds", controller.DeleteLogByIds)

	r.GET("/loginLog/queryLoginLogList", controller.QueryLoginLogList)
	r.POST("/loginLog/deleteLoginLogByIds", controller.DeleteLoginLogByIds)
}
