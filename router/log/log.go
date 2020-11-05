package log

import (
	"github.com/gin-gonic/gin"
	"simple-go/api/log"
)

func LogUrl(r *gin.RouterGroup) {

	r.POST("/syslog/add", log.CreateLog)
	r.POST("/syslog/list", log.GetLogList)
	r.POST("/syslog/update", log.UpdateLog)
	r.POST("/syslog/delete", log.DeleteLogById)

	r.POST("/loginLog/add", log.CreateLoginLog)
	r.POST("/loginLog/list", log.GetLoginLogList)
	r.POST("/loginLog/update", log.UpdateLoginLog)
	r.POST("/loginLog/delete", log.DeleteLoginLogById)
}
