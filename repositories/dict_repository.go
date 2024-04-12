package repositories

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"time"
)

func CreatDict(dto dto.DictDto) error {

	dict := models.SysDict{}
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks
	dict.Sort = dto.Sort
	dict.CreateBy = "liufeihua"
	dict.CreateTime = time.Now()
	dict.LastUpdateBy = "liufeihua"
	dict.LastUpdateTime = time.Now()

	err := models.DB.Create(&dict).Error

	return err
}

func GetDictList(current int, pageSize int) ([]models.SysDict, int) {

	var total = 0
	var dict []models.SysDict
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&dict)

	models.DB.Model(&models.SysDict{}).Count(&total)

	return dict, total
}

func UpdateDict(dto dto.DictDto) error {
	dict := models.SysDict{}
	dict.Id = dto.Id
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Sort = dto.Sort
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks
	dict.LastUpdateBy = "liufeihua"
	dict.LastUpdateTime = time.Now()

	err := models.DB.Model(&dict).Update(&dict).Error

	return err
}

func DeleteDictById(id int64) error {
	err := models.DB.Delete(&models.SysDict{Id: id}).Error
	return err
}
