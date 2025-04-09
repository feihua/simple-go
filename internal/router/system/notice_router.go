package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// NoticeRouter 通知公告相关路由
func NoticeRouter(r *gin.RouterGroup, b *a.NoticeController) {

	r.POST("/system/notice/addNotice", b.CreateNotice)
	r.POST("/system/notice/deleteNoticeByIds", b.DeleteNoticeByIds)
	r.POST("/system/notice/updateNotice", b.UpdateNotice)
	r.POST("/system/notice/updateNoticeStatus", b.UpdateNoticeStatus)
	r.POST("/system/notice/queryNoticeDetail", b.QueryNoticeDetail)
	r.POST("/system/notice/queryNoticeList", b.QueryNoticeList)

}
