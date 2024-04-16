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

// QueryUserRoleList 查询用户角色
func (u UserRoleDao) QueryUserRoleList(roleId int64) []int64 {

	var userRoles []models.UserRole
	u.db.Where("user_id=?", roleId).Find(&userRoles)

	var menuIds []int64

	for _, userRole := range userRoles {
		menuIds = append(menuIds, userRole.UserId)
	}

	return menuIds
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
