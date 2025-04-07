package dao

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
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

	dict := model.Dict{}
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Description = dto.Description
	dict.Remark = dto.Remark
	dict.Sort = dto.Sort

	return d.db.Create(&dict).Error
}

// QueryDictList 查询字典列表
func (d DictDao) QueryDictList(current int, pageSize int) ([]model.Dict, int64) {

	var total int64 = 0
	var dict []model.Dict
	d.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&dict)

	d.db.Model(&model.Dict{}).Count(&total)

	return dict, total
}

func (d DictDao) QueryDictByName(name, typeName string) (*model.Dict, error) {

	var dict model.Dict
	err := d.db.First(&dict, "value = ? and 'type' = ?", name, typeName).Error

	switch {
	case err == nil:
		return &dict, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// UpdateDict 更新字典
func (d DictDao) UpdateDict(dto dto.DictDto) error {
	dict := model.Dict{}
	dict.Id = dto.Id
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Sort = dto.Sort
	dict.Description = dto.Description
	dict.Remark = dto.Remark

	return d.db.Model(&dict).Updates(&dict).Error
}

// DeleteDictByIds 删除字典
func (d DictDao) DeleteDictByIds(ids []int64) error {
	return d.db.Where("id in (?)", ids).Delete(&model.Dict{}).Error
}
