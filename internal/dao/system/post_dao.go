package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
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
	item := a.Post{
		Id:         dto.Id,         // 岗位id
		PostCode:   dto.PostCode,   // 岗位编码
		PostName:   dto.PostName,   // 岗位名称
		Sort:       dto.Sort,       // 显示顺序
		Status:     dto.Status,     // 岗位状态（0：停用，1:正常）
		Remark:     dto.Remark,     // 备注
		CreateBy:   dto.CreateBy,   // 创建者
		CreateTime: dto.CreateTime, // 创建时间
		UpdateBy:   dto.UpdateBy,   // 更新者
		UpdateTime: dto.UpdateTime, // 更新时间
	}

	return b.db.Create(&item).Error
}

// DeletePostByIds 根据id删除岗位信息
func (b PostDao) DeletePostByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.Post{}).Error
}

// UpdatePost 更新岗位信息
func (b PostDao) UpdatePost(dto system.UpdatePostDto) error {

	item := a.Post{
		Id:         dto.Id,         // 岗位id
		PostCode:   dto.PostCode,   // 岗位编码
		PostName:   dto.PostName,   // 岗位名称
		Sort:       dto.Sort,       // 显示顺序
		Status:     dto.Status,     // 岗位状态（0：停用，1:正常）
		Remark:     dto.Remark,     // 备注
		CreateBy:   dto.CreateBy,   // 创建者
		CreateTime: dto.CreateTime, // 创建时间
		UpdateBy:   dto.UpdateBy,   // 更新者
		UpdateTime: dto.UpdateTime, // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdatePostStatus 更新岗位信息状态
func (b PostDao) UpdatePostStatus(dto system.UpdatePostStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryPostDetail 查询岗位信息详情
func (b PostDao) QueryPostDetail(dto system.QueryPostDetailDto) (a.Post, error) {
	var item a.Post
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryPostList 查询岗位信息列表
func (b PostDao) QueryPostList(dto system.QueryPostListDto) ([]a.Post, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.Post
	tx := b.db.Model(&a.Post{})
	if len(dto.PostCode) > 0 {
		tx.Where("post_code like %?%", dto.PostCode) // 岗位编码
	}
	if len(dto.PostName) > 0 {
		tx.Where("post_name like %?%", dto.PostName) // 岗位名称
	}
	if dto.Status != 2 {
		tx.Where("status=?", dto.Status) // 岗位状态（0：停用，1:正常）
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
