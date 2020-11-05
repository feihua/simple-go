package repositories

import (
	"simple-go/models"
	"time"
)

func UpdateUserRole(userId int64,roleId int64) error {

	//先删除
	models.DB.Raw("DELETE FROM sys_user_role WHERE user_id = ?", userId)

	//后添加
	userRole := models.SysUserRole{}
	userRole.UserId=userId
	userRole.RoleId=roleId
	userRole.CreateBy="admin"
	userRole.CreateTime=time.Now()
	userRole.CreateBy="admin"
	userRole.LastUpdateTime=time.Now()

	err := models.DB.Create(&userRole).Error

	return err
}
