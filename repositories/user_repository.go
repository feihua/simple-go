package repositories

import (
	"fmt"
	"simple-go/dto"
	"simple-go/models"
)

func CreateUser(dto dto.UserDto) error {
	user := models.SysUser{}
	//user.Username = dto.Username
	//user.Password = dto.Password

	err := models.DB.Create(&user).Error

	return err
}

func GetUserByUsername(username string) []models.SysUser {
	var sysUser []models.SysUser
	models.DB.Find(&sysUser, models.DB.Where("username = ?", username))

	return sysUser
}

func GetUserList(current int, pageSize int) ([]models.SysUser,int) {

	var total = 0
	var sysUser []models.SysUser
	models.DB.Limit(pageSize).Offset((current-1)*pageSize).Find(&sysUser)

	models.DB.Model(&models.SysUser{}).Count(&total)

	for k, v := range sysUser {
		fmt.Println(k,v)
	}

	return sysUser,total
}

func UpdateUser(userDto dto.UserDto) error {
	sysUser := models.SysUser{}

	err := models.DB.Model(&sysUser).Update(&sysUser).Error

	return err
}

func DeleteUserById(id int64) error {
	err := models.DB.Delete(&models.SysUser{Id: id}).Error
	return err
}
