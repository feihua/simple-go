package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// NoticeRouter 通知公告相关路由
func NoticeRouter(r *gin.RouterGroup, b *a.NoticeController) {

	r.POST("/dept/addNotice", b.CreateNotice)
	r.POST("/dept/deleteNoticeByIds", b.DeleteNoticeByIds)
	r.POST("/dept/updateNotice", b.UpdateNotice)
	r.POST("/dept/updateNoticeStatus", b.UpdateNoticeStatus)
	r.POST("/dept/queryNoticeDetail", b.QueryNoticeDetail)
	r.POST("/dept/queryNoticeList", b.QueryNoticeList)

}
