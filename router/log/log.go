package log

import (
	"github.com/feihua/simple-go/api/log"
	"github.com/gin-gonic/gin"
)

func LogUrl(r *gin.RouterGroup) {

	r.POST("/syslog/add", log.CreateLog)
	r.GET("/syslog/list", log.GetLogList)
	r.POST("/syslog/update", log.UpdateLog)
	r.POST("/syslog/delete", log.DeleteLogById)

	r.POST("/loginLog/add", log.CreateLoginLog)
	r.GET("/loginLog/list", log.GetLoginLogList)
	r.POST("/loginLog/update", log.UpdateLoginLog)
	r.POST("/loginLog/delete", log.DeleteLoginLogById)
}
