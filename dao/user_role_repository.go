package dao

import (
	"github.com/feihua/simple-go/models"
	"time"
)

func UpdateUserRole(userId int64, roleId int64) error {

	//先删除
	models.DB.Where("user_id = ?", userId).Delete(&models.UserRole{})

	//后添加
	userRole := models.UserRole{}
	userRole.UserId = userId
	userRole.RoleId = roleId
	userRole.CreateTime = time.Now()

	err := models.DB.Create(&userRole).Error

	return err
}
