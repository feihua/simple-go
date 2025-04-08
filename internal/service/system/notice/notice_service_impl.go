package notice

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// NoticeServiceImpl 通知公告操作实现
type NoticeServiceImpl struct {
	Dao *system.NoticeDao
}

func NewNoticeServiceImpl(dao *system.NoticeDao) NoticeService {
	return &NoticeServiceImpl{
		Dao: dao,
	}
}

// CreateNotice 添加通知公告
func (s *NoticeServiceImpl) CreateNotice(dto a.AddNoticeDto) error {
	return s.Dao.CreateNotice(dto)
}

// DeleteNoticeByIds 删除通知公告
func (s *NoticeServiceImpl) DeleteNoticeByIds(ids []int64) error {
	return s.Dao.DeleteNoticeByIds(ids)
}

// UpdateNotice 更新通知公告
func (s *NoticeServiceImpl) UpdateNotice(dto a.UpdateNoticeDto) error {
	return s.Dao.UpdateNotice(dto)
}

// UpdateNoticeStatus 更新通知公告状态
func (s *NoticeServiceImpl) UpdateNoticeStatus(dto a.UpdateNoticeStatusDto) error {
	return s.Dao.UpdateNoticeStatus(dto)
}

// QueryNoticeDetail 查询通知公告详情
func (s *NoticeServiceImpl) QueryNoticeDetail(dto a.QueryNoticeDetailDto) (b.Notice, error) {
	return s.Dao.QueryNoticeDetail(dto)
}

// QueryNoticeList 查询通知公告列表
func (s *NoticeServiceImpl) QueryNoticeList(dto a.QueryNoticeListDto) ([]b.Notice, int64) {
	return s.Dao.QueryNoticeList(dto)
}
