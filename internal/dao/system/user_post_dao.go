package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
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
	item := a.UserPost{
		UserId: dto.UserId, // 用户id
		PostId: dto.PostId, // 岗位id
	}

	return b.db.Create(&item).Error
}

// DeleteUserPostByIds 根据id删除用户与岗位关联
func (b UserPostDao) DeleteUserPostByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.UserPost{}).Error
}

// QueryUserPostList 查询用户与岗位关联列表
func (b UserPostDao) QueryUserPostList(dto system.QueryUserPostListDto) ([]a.UserPost, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.UserPost
	tx := b.db.Model(&a.UserPost{})
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
