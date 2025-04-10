package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type UserPostDao struct {
	db *gorm.DB
}

func NewUserPostDao(DB *gorm.DB) *UserPostDao {
	return &UserPostDao{
		db: DB,
	}
}

// CreateUserPost 添加用户与岗位关联
func (b UserPostDao) CreateUserPost(dto system.AddUserPostDto) error {
	item := m.UserPost{
		UserId: dto.UserId, // 用户id
		PostId: dto.PostId, // 岗位id
	}

	return b.db.Create(&item).Error
}

// DeleteUserPostByUserId 根据userId删除用户与岗位关联
func (b UserPostDao) DeleteUserPostByUserId(userId int64) error {
	return b.db.Where("user_id = ?", userId).Delete(&m.UserPost{}).Error
}

// DeleteUserPostByUserIds 根据userIds删除用户与岗位关联
func (b UserPostDao) DeleteUserPostByUserIds(userIds []int64) error {
	return b.db.Where("user_id in (?)", userIds).Delete(&m.UserPost{}).Error
}

// QueryUserPostList 查询用户与岗位关联列表
func (b UserPostDao) QueryUserPostList(dto system.QueryUserPostListDto) ([]m.UserPost, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []m.UserPost
	tx := b.db.Model(&m.UserPost{})
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
