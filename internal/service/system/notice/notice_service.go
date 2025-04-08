package notice

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// NoticeService 通知公告操作接口
type NoticeService interface {
	// CreateNotice 添加通知公告
	CreateNotice(dto b.AddNoticeDto) error
	// DeleteNoticeByIds 删除通知公告
	DeleteNoticeByIds(ids []int64) error
	// UpdateNotice 更新通知公告
	UpdateNotice(dto b.UpdateNoticeDto) error
	// UpdateNoticeStatus 更新通知公告状态
	UpdateNoticeStatus(dto b.UpdateNoticeStatusDto) error
	// QueryNoticeDetail 查询通知公告详情
	QueryNoticeDetail(dto b.QueryNoticeDetailDto) (a.Notice, error)
	// QueryNoticeList 查询通知公告列表
	QueryNoticeList(dto b.QueryNoticeListDto) ([]a.Notice, int64)
}
