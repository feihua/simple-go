package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type NoticeDao struct {
	db *gorm.DB
}

func NewNoticeDao(DB *gorm.DB) *NoticeDao {
	return &NoticeDao{
		db: DB,
	}
}

// CreateNotice 添加通知公告
func (b NoticeDao) CreateNotice(dto system.AddNoticeDto) error {
	item := m.Notice{
		NoticeTitle:   dto.NoticeTitle,   // 公告标题
		NoticeType:    dto.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: dto.NoticeContent, // 公告内容
		Status:        dto.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        dto.Remark,        // 备注
		CreateBy:      dto.CreateBy,      // 创建者

	}

	return b.db.Create(&item).Error
}

// DeleteNoticeByIds 根据id删除通知公告
func (b NoticeDao) DeleteNoticeByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.Notice{}).Error
}

// UpdateNotice 更新通知公告
func (b NoticeDao) UpdateNotice(dto system.UpdateNoticeDto) error {

	item := m.Notice{
		Id:            dto.Id,            // 公告ID
		NoticeTitle:   dto.NoticeTitle,   // 公告标题
		NoticeType:    dto.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: dto.NoticeContent, // 公告内容
		Status:        dto.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        dto.Remark,        // 备注
		CreateBy:      dto.CreateBy,      // 创建者
		CreateTime:    dto.CreateTime,    // 创建时间
		UpdateBy:      dto.UpdateBy,      // 更新者
		UpdateTime:    &dto.UpdateTime,   // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdateNoticeStatus 更新通知公告状态
func (b NoticeDao) UpdateNoticeStatus(dto system.UpdateNoticeStatusDto) error {

	return b.db.Model(&m.Notice{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryNoticeDetail 查询通知公告详情
func (b NoticeDao) QueryNoticeDetail(dto system.QueryNoticeDetailDto) (*m.Notice, error) {
	var item m.Notice
	err := b.db.Where("id", dto.Id).First(&item).Error
	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryNoticeList 查询通知公告列表
func (b NoticeDao) QueryNoticeList(dto system.QueryNoticeListDto) ([]*m.Notice, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.Notice
	tx := b.db.Model(&m.Notice{})
	if len(dto.NoticeTitle) > 0 {
		tx.Where("notice_title like %?%", dto.NoticeTitle) // 公告标题
	}
	if dto.NoticeType != nil {
		tx.Where("notice_type=?", dto.NoticeType) // 公告类型（1:通知,2:公告）
	}

	if dto.Status != nil {
		tx.Where("status=?", dto.Status) // 公告状态（0:关闭,1:正常 ）
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	err := tx.Count(&total).Error
	return list, total, err
}

// QueryNoticeById 根据id查询通知公告
func (b NoticeDao) QueryNoticeById(id int64) (*m.Notice, error) {
	var item m.Notice
	err := b.db.Where("id = ?", id).Scan(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// IsExistTitle 根据title查询通知公告
func (b NoticeDao) IsExistTitle(title string) (*m.Notice, error) {
	var item m.Notice
	err := b.db.Model(&m.Notice{}).Where("notice_title = ?", title).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
