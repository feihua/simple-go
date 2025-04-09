package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// OperateLogRouter 操作日志记录相关路由
func OperateLogRouter(r *gin.RouterGroup, b *a.OperateLogController) {

	r.POST("/dept/addOperateLog", b.CreateOperateLog)
	r.POST("/dept/deleteOperateLogByIds", b.DeleteOperateLogByIds)
	r.POST("/dept/updateOperateLog", b.UpdateOperateLog)
	r.POST("/dept/updateOperateLogStatus", b.UpdateOperateLogStatus)
	r.POST("/dept/queryOperateLogDetail", b.QueryOperateLogDetail)
	r.POST("/dept/queryOperateLogList", b.QueryOperateLogList)

}
