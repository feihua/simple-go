package repositories

import (
	"fmt"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateUser(dto dto.UserDto) error {

	user := models.User{
		Mobile:   dto.Mobile,
		UserName: dto.UserName,
		Password: dto.Password,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := models.DB.Create(&user).Error

	return err
}

func QueryUserByUsername(username string) []models.User {
	var sysUser []models.User

	models.DB.Find(&sysUser, models.DB.Where("username = ?", username))

	return sysUser
}

func QueryUserList(current int, pageSize int) ([]models.User, int) {

	var total = 0
	var sysUser []models.User
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&sysUser)

	models.DB.Model(&models.User{}).Count(&total)

	for k, v := range sysUser {
		fmt.Println(k, v)
	}

	return sysUser, total
}

func UpdateUser(dto dto.UserDto) error {

	user := models.User{
		Id:       dto.Id,
		Mobile:   dto.Mobile,
		UserName: dto.UserName,
		Password: dto.Password,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := models.DB.Model(&user).Update(&user).Error

	return err
}

func DeleteUserById(id int64) error {
	err := models.DB.Delete(&models.User{Id: id}).Error
	return err
}
