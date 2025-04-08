package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
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
	item := a.Notice{
		Id:            dto.Id,            // 公告ID
		NoticeTitle:   dto.NoticeTitle,   // 公告标题
		NoticeType:    dto.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: dto.NoticeContent, // 公告内容
		Status:        dto.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        dto.Remark,        // 备注
		CreateBy:      dto.CreateBy,      // 创建者
		CreateTime:    dto.CreateTime,    // 创建时间
		UpdateBy:      dto.UpdateBy,      // 更新者
		UpdateTime:    dto.UpdateTime,    // 更新时间
	}

	return b.db.Create(&item).Error
}

// DeleteNoticeByIds 根据id删除通知公告
func (b NoticeDao) DeleteNoticeByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.Notice{}).Error
}

// UpdateNotice 更新通知公告
func (b NoticeDao) UpdateNotice(dto system.UpdateNoticeDto) error {

	item := a.Notice{
		Id:            dto.Id,            // 公告ID
		NoticeTitle:   dto.NoticeTitle,   // 公告标题
		NoticeType:    dto.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: dto.NoticeContent, // 公告内容
		Status:        dto.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        dto.Remark,        // 备注
		CreateBy:      dto.CreateBy,      // 创建者
		CreateTime:    dto.CreateTime,    // 创建时间
		UpdateBy:      dto.UpdateBy,      // 更新者
		UpdateTime:    dto.UpdateTime,    // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdateNoticeStatus 更新通知公告状态
func (b NoticeDao) UpdateNoticeStatus(dto system.UpdateNoticeStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryNoticeDetail 查询通知公告详情
func (b NoticeDao) QueryNoticeDetail(dto system.QueryNoticeDetailDto) (a.Notice, error) {
	var item a.Notice
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryNoticeList 查询通知公告列表
func (b NoticeDao) QueryNoticeList(dto system.QueryNoticeListDto) ([]a.Notice, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.Notice
	tx := b.db.Model(&a.Notice{})
	if len(dto.NoticeTitle) > 0 {
		tx.Where("notice_title like %?%", dto.NoticeTitle) // 公告标题
	}
	if dto.NoticeType != 2 {
		tx.Where("notice_type=?", dto.NoticeType) // 公告类型（1:通知,2:公告）
	}
	if len(dto.NoticeContent) > 0 {
		tx.Where("notice_content like %?%", dto.NoticeContent) // 公告内容
	}
	if dto.Status != 2 {
		tx.Where("status=?", dto.Status) // 公告状态（0:关闭,1:正常 ）
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
