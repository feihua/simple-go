package dao

import (
	"github.com/feihua/simple-go/models"
	"gorm.io/gorm"
)

type UserRoleDao struct {
	db *gorm.DB
}

func NewUserRoleDao(DB *gorm.DB) *UserRoleDao {
	return &UserRoleDao{db: DB}
}

// IsAdministrator 根据用户id判断是否是管理员
func (u UserRoleDao) IsAdministrator(userId int64) bool {

	var count int64
	//1是预留超级管理员角色的id
	u.db.Model(&models.UserRole{}).Where("user_id= ? and role_id = 1", userId).Count(&count)

	return count == 1
}

// QueryUserRoleList 根据userId查询用户角色的roleId
func (u UserRoleDao) QueryUserRoleList(userId int64) ([]int64, error) {

	var userRoles []models.UserRole
	err := u.db.Where("user_id = ?", userId).Find(&userRoles).Error

	var roleIds []int64

	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	return roleIds, err
}

// UpdateUserRoleList 更新用户与角色关糸
func (u UserRoleDao) UpdateUserRoleList(userId int64, userRoles []models.UserRole) error {

	//先删除
	u.db.Where("user_id = ?", userId).Delete(&models.UserRole{})

	return u.db.Model(&models.UserRole{}).CreateInBatches(&userRoles, len(userRoles)).Error
}
