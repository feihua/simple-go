package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type DictTypeDao struct {
	db *gorm.DB
}

func NewDictTypeDao(DB *gorm.DB) *DictTypeDao {
	return &DictTypeDao{
		db: DB,
	}
}

// CreateDictType 添加字典类型
func (b DictTypeDao) CreateDictType(dto system.AddDictTypeDto) error {
	item := m.DictType{
		DictName: dto.DictName, // 字典名称
		DictType: dto.DictType, // 字典类型
		Status:   dto.Status,   // 状态（0：停用，1:正常）
		Remark:   dto.Remark,   // 备注
		CreateBy: dto.CreateBy, // 创建者

	}

	return b.db.Create(&item).Error
}

// DeleteDictTypeByIds 根据id删除字典类型
func (b DictTypeDao) DeleteDictTypeByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.DictType{}).Error
}

// UpdateDictType 更新字典类型
func (b DictTypeDao) UpdateDictType(dto system.UpdateDictTypeDto) error {

	item := m.DictType{
		Id:         dto.Id,          // 字典主键
		DictName:   dto.DictName,    // 字典名称
		DictType:   dto.DictType,    // 字典类型
		Status:     dto.Status,      // 状态（0：停用，1:正常）
		Remark:     dto.Remark,      // 备注
		CreateBy:   dto.CreateBy,    // 创建者
		CreateTime: dto.CreateTime,  // 创建时间
		UpdateBy:   dto.UpdateBy,    // 更新者
		UpdateTime: &dto.UpdateTime, // 更新时间
	}

	return b.db.Save(&item).Error
}

// UpdateDictTypeStatus 更新字典类型状态
func (b DictTypeDao) UpdateDictTypeStatus(dto system.UpdateDictTypeStatusDto) error {

	return b.db.Model(&m.DictType{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryDictTypeDetail 查询字典类型详情
func (b DictTypeDao) QueryDictTypeDetail(dto system.QueryDictTypeDetailDto) (*m.DictType, error) {
	var item m.DictType
	err := b.db.Where("id", dto.Id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryDictTypeList 查询字典类型列表
func (b DictTypeDao) QueryDictTypeList(dto system.QueryDictTypeListDto) ([]*m.DictType, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.DictType
	tx := b.db.Model(&m.DictType{})
	if len(dto.DictName) > 0 {
		tx.Where("dict_name like %?%", dto.DictName) // 字典名称
	}
	if len(dto.DictType) > 0 {
		tx.Where("dict_type like %?%", dto.DictType) // 字典类型
	}
	if dto.Status != nil {
		tx.Where("status=?", dto.Status) // 状态（0：停用，1:正常）
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	err := tx.Count(&total).Error
	return list, total, err
}

// QueryDictTypeById 根据id查询字典类型
func (b DictTypeDao) QueryDictTypeById(id int64) (*m.DictType, error) {
	var item m.DictType
	err := b.db.Where("id = ?", id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryDictTypeByName 根据name查询字典类型
func (b DictTypeDao) QueryDictTypeByName(name string) (*m.DictType, error) {
	var item m.DictType
	err := b.db.Where("dict_name = ?", name).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryDictTypeByType 根据type查询字典类型
func (b DictTypeDao) QueryDictTypeByType(dictType string) (*m.DictType, error) {
	var item m.DictType
	err := b.db.Where("dict_type = ?", dictType).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
