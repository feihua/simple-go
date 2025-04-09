package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
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
	item := a.DictType{
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
	return b.db.Where("id in (?)", ids).Delete(&a.DictType{}).Error
}

// UpdateDictType 更新字典类型
func (b DictTypeDao) UpdateDictType(dto system.UpdateDictTypeDto) error {

	item := a.DictType{
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

	return b.db.Updates(&item).Error
}

// UpdateDictTypeStatus 更新字典类型状态
func (b DictTypeDao) UpdateDictTypeStatus(dto system.UpdateDictTypeStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryDictTypeDetail 查询字典类型详情
func (b DictTypeDao) QueryDictTypeDetail(dto system.QueryDictTypeDetailDto) (a.DictType, error) {
	var item a.DictType
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryDictTypeList 查询字典类型列表
func (b DictTypeDao) QueryDictTypeList(dto system.QueryDictTypeListDto) ([]a.DictType, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.DictType
	tx := b.db.Model(&a.DictType{})
	if len(dto.DictName) > 0 {
		tx.Where("dict_name like %?%", dto.DictName) // 字典名称
	}
	if len(dto.DictType) > 0 {
		tx.Where("dict_type like %?%", dto.DictType) // 字典类型
	}
	if dto.Status != 2 {
		tx.Where("status=?", dto.Status) // 状态（0：停用，1:正常）
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
