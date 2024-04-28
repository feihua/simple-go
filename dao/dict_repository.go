package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"gorm.io/gorm"
)

type DictDao struct {
	db *gorm.DB
}

func NewDictDao(DB *gorm.DB) *DictDao {
	return &DictDao{db: DB}
}

// CreateDict 创建字典
func (d DictDao) CreateDict(dto dto.DictDto) error {

	dict := models.Dict{}
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks
	dict.Sort = dto.Sort

	return d.db.Create(&dict).Error
}

// QueryDictList 查询字典列表
func (d DictDao) QueryDictList(current int, pageSize int) ([]models.Dict, int64) {

	var total int64 = 0
	var dict []models.Dict
	d.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&dict)

	d.db.Model(&models.Dict{}).Count(&total)

	return dict, total
}

func (d DictDao) QueryDictByName(name, typeName string) (*models.Dict, error) {

	var dict models.Dict
	err := d.db.First(&dict, "value = ? and 'type' = ?", name, typeName).Error

	return &dict, err
}

// UpdateDict 更新字典
func (d DictDao) UpdateDict(dto dto.DictDto) error {
	dict := models.Dict{}
	dict.Id = dto.Id
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Sort = dto.Sort
	dict.Description = dto.Description
	dict.Remarks = dto.Remarks

	return d.db.Model(&dict).Updates(&dict).Error
}

// DeleteDictByIds 删除字典
func (d DictDao) DeleteDictByIds(ids []int64) error {
	return d.db.Where("id in (?)", ids).Delete(&models.Dict{}).Error
}
