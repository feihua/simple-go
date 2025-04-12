package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// OperateLogRouter 操作日志记录相关路由
func OperateLogRouter(r *gin.RouterGroup, b *a.OperateLogController) {

	r.POST("/system/operateLog/deleteOperateLog", b.DeleteOperateLogByIds)
	r.POST("/system/operateLog/queryOperateLogDetail", b.QueryOperateLogDetail)
	r.POST("/system/operateLog/queryOperateLogList", b.QueryOperateLogList)

}
