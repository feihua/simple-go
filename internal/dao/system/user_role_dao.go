package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type UserRoleDao struct {
	db *gorm.DB
}

func NewUserRoleDao(DB *gorm.DB) *UserRoleDao {
	return &UserRoleDao{
		db: DB,
	}
}

// CreateUserRole 添加角色用户关联
func (b UserRoleDao) CreateUserRole(dto system.AddUserRoleDto) error {
	item := m.UserRole{
		UserId: dto.UserId, // 用户ID
		RoleId: dto.RoleId, // 角色ID
	}

	return b.db.Create(&item).Error
}

// DeleteUserRoleByUserId 根据userId删除角色用户关联
func (b UserRoleDao) DeleteUserRoleByUserId(userId int64) error {
	return b.db.Where("user_id = ?", userId).Delete(&m.UserRole{}).Error
}

// DeleteUserRoleByRoleId 根据roleId删除角色用户关联
func (b UserRoleDao) DeleteUserRoleByRoleId(roleId int64) error {
	return b.db.Where("role_id = ?", roleId).Delete(&m.UserRole{}).Error
}

// DeleteUserRoleByUserIds 根据userIds删除角色用户关联
func (b UserRoleDao) DeleteUserRoleByUserIds(userIds []int64) error {
	return b.db.Where("user_id in (?)", userIds).Delete(&m.UserRole{}).Error
}

// QueryUserRoleIds 查询角色ids
func (b UserRoleDao) QueryUserRoleIds(userId int64) ([]int64, error) {

	var ids []int64
	err := b.db.Model(&m.UserRole{}).Select("role_id").Where("user_id=?", userId).Scan(&ids).Error
	return ids, err

}

// IsAdministrator 根据用户id判断是否是管理员
func (u UserRoleDao) IsAdministrator(userId int64) bool {

	var count int64
	// 1是预留超级管理员角色的id
	u.db.Model(&m.UserRole{}).Where("user_id= ? and role_id = 1", userId).Count(&count)

	return count > 0
}
