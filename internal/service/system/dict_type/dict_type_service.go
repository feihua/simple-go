package dict_type

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// DictTypeService 字典类型操作接口
type DictTypeService interface {
	// CreateDictType 添加字典类型
	CreateDictType(dto d.AddDictTypeDto) error
	// DeleteDictTypeByIds 删除字典类型
	DeleteDictTypeByIds(ids []int64) error
	// UpdateDictType 更新字典类型
	UpdateDictType(dto d.UpdateDictTypeDto) error
	// UpdateDictTypeStatus 更新字典类型状态
	UpdateDictTypeStatus(dto d.UpdateDictTypeStatusDto) error
	// QueryDictTypeDetail 查询字典类型详情
	QueryDictTypeDetail(dto d.QueryDictTypeDetailDto) (*d.QueryDictTypeListDtoResp, error)
	// QueryDictTypeById 根据id查询字典类型详情
	QueryDictTypeById(id int64) (*d.QueryDictTypeListDtoResp, error)
	// QueryDictTypeList 查询字典类型列表
	QueryDictTypeList(dto d.QueryDictTypeListDto) ([]*d.QueryDictTypeListDtoResp, int64, error)
}
