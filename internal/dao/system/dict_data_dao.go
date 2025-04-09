package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type DictDataDao struct {
	db *gorm.DB
}

func NewDictDataDao(DB *gorm.DB) *DictDataDao {
	return &DictDataDao{
		db: DB,
	}
}

// CreateDictData 添加字典数据
func (b DictDataDao) CreateDictData(dto system.AddDictDataDto) error {
	item := a.DictData{
		DictSort:  dto.DictSort,  // 字典排序
		DictLabel: dto.DictLabel, // 字典标签
		DictValue: dto.DictValue, // 字典键值
		DictType:  dto.DictType,  // 字典类型
		CssClass:  dto.CssClass,  // 样式属性（其他样式扩展）
		ListClass: dto.ListClass, // 表格回显样式
		IsDefault: dto.IsDefault, // 是否默认（Y是 N否）
		Status:    dto.Status,    // 状态（0：停用，1:正常）
		Remark:    dto.Remark,    // 备注
		CreateBy:  dto.CreateBy,  // 创建者
	}

	return b.db.Create(&item).Error
}

// DeleteDictDataByIds 根据id删除字典数据
func (b DictDataDao) DeleteDictDataByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.DictData{}).Error
}

// UpdateDictData 更新字典数据
func (b DictDataDao) UpdateDictData(dto system.UpdateDictDataDto) error {

	item := a.DictData{
		Id:         dto.Id,          // 字典编码
		DictSort:   dto.DictSort,    // 字典排序
		DictLabel:  dto.DictLabel,   // 字典标签
		DictValue:  dto.DictValue,   // 字典键值
		DictType:   dto.DictType,    // 字典类型
		CssClass:   dto.CssClass,    // 样式属性（其他样式扩展）
		ListClass:  dto.ListClass,   // 表格回显样式
		IsDefault:  dto.IsDefault,   // 是否默认（Y是 N否）
		Status:     dto.Status,      // 状态（0：停用，1:正常）
		Remark:     dto.Remark,      // 备注
		CreateBy:   dto.CreateBy,    // 创建者
		CreateTime: dto.CreateTime,  // 创建时间
		UpdateBy:   dto.UpdateBy,    // 更新者
		UpdateTime: &dto.UpdateTime, // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdateDictDataStatus 更新字典数据状态
func (b DictDataDao) UpdateDictDataStatus(dto system.UpdateDictDataStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryDictDataDetail 查询字典数据详情
func (b DictDataDao) QueryDictDataDetail(dto system.QueryDictDataDetailDto) (a.DictData, error) {
	var item a.DictData
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryDictDataList 查询字典数据列表
func (b DictDataDao) QueryDictDataList(dto system.QueryDictDataListDto) ([]a.DictData, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.DictData
	tx := b.db.Model(&a.DictData{})
	if len(dto.DictLabel) > 0 {
		tx.Where("dict_label like %?%", dto.DictLabel) // 字典标签
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
