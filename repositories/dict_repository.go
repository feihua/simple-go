package repositories

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreatDict(dto dto.DictDto) error {

	dict := models.Dict{}
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks
	dict.Sort = dto.Sort

	err := models.DB.Create(&dict).Error

	return err
}

func QueryDictList(current int, pageSize int) ([]models.Dict, int) {

	var total = 0
	var dict []models.Dict
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&dict)

	models.DB.Model(&models.Dict{}).Count(&total)

	return dict, total
}

func UpdateDict(dto dto.DictDto) error {
	dict := models.Dict{}
	dict.Id = dto.Id
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Sort = dto.Sort
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks

	err := models.DB.Model(&dict).Update(&dict).Error

	return err
}

func DeleteDictById(id int64) error {
	err := models.DB.Delete(&models.Dict{Id: id}).Error
	return err
}
