package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// CreateDict 创建字典
func CreateDict(dto dto.DictDto) error {

	dict := models.Dict{}
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks
	dict.Sort = dto.Sort

	return models.DB.Create(&dict).Error
}

// QueryDictList 查询字典列表
func QueryDictList(current int, pageSize int) ([]models.Dict, int) {

	var total = 0
	var dict []models.Dict
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&dict)

	models.DB.Model(&models.Dict{}).Count(&total)

	return dict, total
}

func QueryDictByName(name, typeName string) (*models.Dict, error) {

	var dict models.Dict
	err := models.DB.First(&dict, "value = ? and 'type' = ?", name, typeName).Error

	return &dict, err
}

// UpdateDict 更新字典
func UpdateDict(dto dto.DictDto) error {
	dict := models.Dict{}
	dict.Id = dto.Id
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Sort = dto.Sort
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks

	return models.DB.Model(&dict).Update(&dict).Error
}

// DeleteDictByIds 删除字典
func DeleteDictByIds(ids []int64) error {
	return models.DB.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}
