package notice

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// NoticeService 通知公告操作接口
type NoticeService interface {
	// CreateNotice 添加通知公告
	CreateNotice(dto d.AddNoticeDto) error
	// DeleteNoticeByIds 删除通知公告
	DeleteNoticeByIds(ids []int64) error
	// UpdateNotice 更新通知公告
	UpdateNotice(dto d.UpdateNoticeDto) error
	// UpdateNoticeStatus 更新通知公告状态
	UpdateNoticeStatus(dto d.UpdateNoticeStatusDto) error
	// QueryNoticeDetail 查询通知公告详情
	QueryNoticeDetail(dto d.QueryNoticeDetailDto) (*d.QueryNoticeListDtoResp, error)
	// QueryNoticeById 根据id查询通知公告详情
	QueryNoticeById(id int64) (*d.QueryNoticeListDtoResp, error)
	// QueryNoticeList 查询通知公告列表
	QueryNoticeList(dto d.QueryNoticeListDto) ([]*d.QueryNoticeListDtoResp, int64, error)
}
