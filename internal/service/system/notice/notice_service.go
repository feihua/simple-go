package notice

import (
	dto "github.com/feihua/simple-go/internal/dto/system"
)

// NoticeService 通知公告操作接口
type NoticeService interface {
	// CreateNotice 添加通知公告
	CreateNotice(req dto.AddNoticeDto) error
	// DeleteNoticeByIds 删除通知公告
	DeleteNoticeByIds(ids []int64) error
	// UpdateNotice 更新通知公告
	UpdateNotice(req dto.UpdateNoticeDto) error
	// UpdateNoticeStatus 更新通知公告状态
	UpdateNoticeStatus(req dto.UpdateNoticeStatusDto) error
	// QueryNoticeDetail 查询通知公告详情
	QueryNoticeDetail(req dto.QueryNoticeDetailDto) (*dto.QueryNoticeListDtoResp, error)
	// QueryNoticeById 根据id查询通知公告详情
	QueryNoticeById(id int64) (*dto.QueryNoticeListDtoResp, error)
	// QueryNoticeList 查询通知公告列表
	QueryNoticeList(req dto.QueryNoticeListDto) ([]*dto.QueryNoticeListDtoResp, int64, error)
}
