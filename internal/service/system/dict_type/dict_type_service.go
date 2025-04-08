package dict_type

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// DictTypeService 字典类型操作接口
type DictTypeService interface {
	// CreateDictType 添加字典类型
	CreateDictType(dto b.AddDictTypeDto) error
	// DeleteDictTypeByIds 删除字典类型
	DeleteDictTypeByIds(ids []int64) error
	// UpdateDictType 更新字典类型
	UpdateDictType(dto b.UpdateDictTypeDto) error
	// UpdateDictTypeStatus 更新字典类型状态
	UpdateDictTypeStatus(dto b.UpdateDictTypeStatusDto) error
	// QueryDictTypeDetail 查询字典类型详情
	QueryDictTypeDetail(dto b.QueryDictTypeDetailDto) (a.DictType, error)
	// QueryDictTypeList 查询字典类型列表
	QueryDictTypeList(dto b.QueryDictTypeListDto) ([]a.DictType, int64)
}
