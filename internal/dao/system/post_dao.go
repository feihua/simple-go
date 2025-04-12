package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type PostDao struct {
	db *gorm.DB
}

func NewPostDao(DB *gorm.DB) *PostDao {
	return &PostDao{
		db: DB,
	}
}

// CreatePost 添加岗位信息
func (b PostDao) CreatePost(dto system.AddPostDto) error {
	item := m.Post{
		PostCode: dto.PostCode, // 岗位编码
		PostName: dto.PostName, // 岗位名称
		Sort:     dto.Sort,     // 显示顺序
		Status:   dto.Status,   // 岗位状态（0：停用，1:正常）
		Remark:   dto.Remark,   // 备注
		CreateBy: dto.CreateBy, // 创建者

	}

	return b.db.Create(&item).Error
}

// DeletePostByIds 根据id删除岗位信息
func (b PostDao) DeletePostByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.Post{}).Error
}

// UpdatePost 更新岗位信息
func (b PostDao) UpdatePost(dto system.UpdatePostDto) error {

	item := m.Post{
		Id:         dto.Id,          // 岗位id
		PostCode:   dto.PostCode,    // 岗位编码
		PostName:   dto.PostName,    // 岗位名称
		Sort:       dto.Sort,        // 显示顺序
		Status:     dto.Status,      // 岗位状态（0：停用，1:正常）
		Remark:     dto.Remark,      // 备注
		CreateBy:   dto.CreateBy,    // 创建者
		CreateTime: dto.CreateTime,  // 创建时间
		UpdateBy:   dto.UpdateBy,    // 更新者
		UpdateTime: &dto.UpdateTime, // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdatePostStatus 更新岗位信息状态
func (b PostDao) UpdatePostStatus(dto system.UpdatePostStatusDto) error {

	return b.db.Model(&m.Post{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryPostDetail 查询岗位信息详情
func (b PostDao) QueryPostDetail(dto system.QueryPostDetailDto) (*m.Post, error) {
	var item m.Post
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

// QueryPostList 查询岗位信息列表
func (b PostDao) QueryPostList(dto system.QueryPostListDto) ([]*m.Post, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.Post
	tx := b.db.Model(&m.Post{})
	if len(dto.PostCode) > 0 {
		tx.Where("post_code like %?%", dto.PostCode) // 岗位编码
	}
	if len(dto.PostName) > 0 {
		tx.Where("post_name like %?%", dto.PostName) // 岗位名称
	}
	if dto.Status != nil {
		tx.Where("status=?", dto.Status) // 岗位状态（0：停用，1:正常）
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	err := tx.Count(&total).Error
	return list, total, err
}

// QueryPostById 根据id查询岗位
func (b PostDao) QueryPostById(id int64) (*m.Post, error) {
	var item m.Post
	err := b.db.Where("id = ?", id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryPostByCode 根据岗位编码查询岗位
func (b PostDao) QueryPostByCode(code string) (*m.Post, error) {
	var item m.Post
	err := b.db.Where("post_code = ?", code).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryPostByName 根据岗位名称查询岗位
func (b PostDao) QueryPostByName(name string) (*m.Post, error) {
	var item m.Post
	err := b.db.Where("post_name = ?", name).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
