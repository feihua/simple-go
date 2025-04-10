package notice

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"time"
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
func (s *NoticeServiceImpl) CreateNotice(dto d.AddNoticeDto) error {
	return s.Dao.CreateNotice(dto)
}

// DeleteNoticeByIds 删除通知公告
func (s *NoticeServiceImpl) DeleteNoticeByIds(ids []int64) error {
	return s.Dao.DeleteNoticeByIds(ids)
}

// UpdateNotice 更新通知公告
func (s *NoticeServiceImpl) UpdateNotice(dto d.UpdateNoticeDto) error {
	item, err := s.Dao.QueryNoticeById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("通知公告不存在")
	}

	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateNotice(dto)
}

// UpdateNoticeStatus 更新通知公告状态
func (s *NoticeServiceImpl) UpdateNoticeStatus(dto d.UpdateNoticeStatusDto) error {
	return s.Dao.UpdateNoticeStatus(dto)
}

// QueryNoticeDetail 查询通知公告详情
func (s *NoticeServiceImpl) QueryNoticeDetail(dto d.QueryNoticeDetailDto) (*d.QueryNoticeListDtoResp, error) {
	item, err := s.Dao.QueryNoticeDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("通知公告不存在")
	}

	return &d.QueryNoticeListDtoResp{
		Id:            item.Id,                             // 公告ID
		NoticeTitle:   item.NoticeTitle,                    // 公告标题
		NoticeType:    item.NoticeType,                     // 公告类型（1:通知,2:公告）
		NoticeContent: item.NoticeContent,                  // 公告内容
		Status:        item.Status,                         // 公告状态（0:关闭,1:正常 ）
		Remark:        item.Remark,                         // 备注
		CreateBy:      item.CreateBy,                       // 创建者
		CreateTime:    utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:      item.UpdateBy,                       // 更新者
		UpdateTime:    utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryNoticeById 根据id查询通知公告详情
func (s *NoticeServiceImpl) QueryNoticeById(id int64) (*d.QueryNoticeListDtoResp, error) {
	item, err := s.Dao.QueryNoticeById(id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("通知公告不存在")
	}

	return &d.QueryNoticeListDtoResp{
		Id:            item.Id,                             // 公告ID
		NoticeTitle:   item.NoticeTitle,                    // 公告标题
		NoticeType:    item.NoticeType,                     // 公告类型（1:通知,2:公告）
		NoticeContent: item.NoticeContent,                  // 公告内容
		Status:        item.Status,                         // 公告状态（0:关闭,1:正常 ）
		Remark:        item.Remark,                         // 备注
		CreateBy:      item.CreateBy,                       // 创建者
		CreateTime:    utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:      item.UpdateBy,                       // 更新者
		UpdateTime:    utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryNoticeList 查询通知公告列表
func (s *NoticeServiceImpl) QueryNoticeList(dto d.QueryNoticeListDto) ([]*d.QueryNoticeListDtoResp, int64, error) {
	result, i, err := s.Dao.QueryNoticeList(dto)

	if err != nil {
		return nil, 0, err
	}

	var list []*d.QueryNoticeListDtoResp

	for _, item := range result {
		resp := &d.QueryNoticeListDtoResp{
			Id:            item.Id,                             // 公告ID
			NoticeTitle:   item.NoticeTitle,                    // 公告标题
			NoticeType:    item.NoticeType,                     // 公告类型（1:通知,2:公告）
			NoticeContent: item.NoticeContent,                  // 公告内容
			Status:        item.Status,                         // 公告状态（0:关闭,1:正常 ）
			Remark:        item.Remark,                         // 备注
			CreateBy:      item.CreateBy,                       // 创建者
			CreateTime:    utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:      item.UpdateBy,                       // 更新者
			UpdateTime:    utils.TimeToString(item.UpdateTime), // 更新时间
		}

		list = append(list, resp)
	}

	return list, i, nil
}
