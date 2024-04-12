package repositories

import (
	"fmt"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"time"
)

func CreateUser(dto dto.UserDto) error {

	user := models.SysUser{}
	user.Name = dto.Name
	user.NickName = dto.NickName
	user.Mobile = dto.Mobile
	user.Email = dto.Email
	user.Status = 1
	user.Password = "123456"
	user.CreateBy = "liufeihua"
	user.CreateTime = time.Now()
	user.LastUpdateBy = "liufeihua"
	user.LastUpdateTime = time.Now()

	err := models.DB.Create(&user).Error

	return err
}

func GetUserByUsername(username string) []models.SysUser {
	var sysUser []models.SysUser

	models.DB.Find(&sysUser, models.DB.Where("username = ?", username))

	return sysUser
}

func GetUserList(current int, pageSize int) ([]models.SysUser, int) {

	var total = 0
	var sysUser []models.SysUser
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&sysUser)

	models.DB.Model(&models.SysUser{}).Count(&total)

	for k, v := range sysUser {
		fmt.Println(k, v)
	}

	return sysUser, total
}

func UpdateUser(dto dto.UserDto) error {

	user := models.SysUser{}
	user.Id = dto.Id
	user.Name = dto.Name
	user.NickName = dto.NickName
	user.Mobile = dto.Mobile
	user.Email = dto.Email
	user.Password = "123456"
	user.LastUpdateBy = "liufeihua"
	user.LastUpdateTime = time.Now()

	err := models.DB.Model(&user).Update(&user).Error

	return err
}

func DeleteUserById(id int64) error {
	err := models.DB.Delete(&models.SysUser{Id: id}).Error
	return err
}
