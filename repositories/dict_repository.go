package repositories

import (
	"simple-go/dto"
	"simple-go/models"
)

func CreatDict(dto dto.DictDto) error {
	dict := models.SysDict{}

	err := models.DB.Create(&dict).Error

	return err
}

func GetDictList(current int, pageSize int) ([]models.SysDict,int) {

	var total = 0
	var dict []models.SysDict
	models.DB.Limit(pageSize).Offset((current-1)*pageSize).Find(&dict)

	models.DB.Model(&models.SysDict{}).Count(&total)

	return dict,total
}

func UpdateDict(userDto dto.DictDto) error {
	dict := models.SysDict{}
	//if userDto.Username != "" {
	//	dict.Username = userDto.Username
	//}
	//
	//if userDto.Password != "" {
	//	dict.Password = userDto.Password
	//}

	err := models.DB.Model(&dict).Update(&dict).Error

	return err
}

func DeleteDictById(id int64) error {
	err := models.DB.Delete(&models.SysDict{Id: id}).Error
	return err
}
