package log

import (
	"github.com/feihua/simple-go/controller/log"
	"github.com/gin-gonic/gin"
)

func LogUrl(r *gin.RouterGroup) {

	r.POST("/syslog/add", log.CreateLog)
	r.GET("/syslog/list", log.QueryLogList)
	r.POST("/syslog/delete", log.DeleteLogById)

	r.POST("/loginLog/add", log.CreateLoginLog)
	r.GET("/loginLog/list", log.QueryLoginLogList)
	r.POST("/loginLog/delete", log.DeleteLoginLogById)
}
