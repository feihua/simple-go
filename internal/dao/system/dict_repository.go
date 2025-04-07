package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type DictDao struct {
	db *gorm.DB
}

func NewDictDao(DB *gorm.DB) *DictDao {
	return &DictDao{db: DB}
}

// CreateDict 创建字典
func (d DictDao) CreateDict(dto system.DictDto) error {

	dict := system2.Dict{}
	dict.Value = dto.Value
	dict.Label = dto.Label
	dict.Type = dto.Type
	dict.Description = dto.Description
	dict.Remark = dto.Remark
	dict.Sort = dto.Sort

	return d.db.Create(&dict).Error
}

// QueryDictList 查询字典列表
func (d DictDao) QueryDictList(current int, pageSize int) ([]system2.Dict, int64) {

	var total int64 = 0
	var dict []system2.Dict
	d.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&dict)

	d.db.Model(&system2.Dict{}).Count(&total)

	return dict, total
}

func (d DictDao) QueryDictByName(name, typeName string) (*system2.Dict, error) {

	var dict system2.Dict
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
func (d DictDao) UpdateDict(dto system.DictDto) error {
	dict := system2.Dict{}
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
	return d.db.Where("id in (?)", ids).Delete(&system2.Dict{}).Error
}
