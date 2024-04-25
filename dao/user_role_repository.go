package dao

import (
	"github.com/feihua/simple-go/models"
	"github.com/jinzhu/gorm"
	"time"
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

// QueryUserRoleList 查询用户角色
func (u UserRoleDao) QueryUserRoleList(userId int64) []int64 {

	var userRoles []models.UserRole
	u.db.Where("user_id = ?", userId).Find(&userRoles)

	var roleIds []int64

	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	return roleIds
}

// UpdateUserRoleList 更新用户与角色关糸
func (u UserRoleDao) UpdateUserRoleList(userId int64, roleId int64) error {

	//先删除
	u.db.Where("user_id = ?", userId).Delete(&models.UserRole{})

	//后添加
	userRole := models.UserRole{}
	userRole.UserId = userId
	userRole.RoleId = roleId
	userRole.CreateTime = time.Now()

	err := u.db.Create(&userRole).Error

	return err
}
