package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// OperateLogRouter 操作日志记录相关路由
func OperateLogRouter(r *gin.RouterGroup, b *a.OperateLogController) {

	r.POST("/system/operateLog/addOperateLog", b.CreateOperateLog)
	r.POST("/system/operateLog/deleteOperateLogByIds", b.DeleteOperateLogByIds)
	r.POST("/system/operateLog/updateOperateLog", b.UpdateOperateLog)
	r.POST("/system/operateLog/updateOperateLogStatus", b.UpdateOperateLogStatus)
	r.POST("/system/operateLog/queryOperateLogDetail", b.QueryOperateLogDetail)
	r.POST("/system/operateLog/queryOperateLogList", b.QueryOperateLogList)

}
