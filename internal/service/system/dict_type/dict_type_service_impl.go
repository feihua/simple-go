package dict_type

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// DictTypeServiceImpl 字典类型操作实现
type DictTypeServiceImpl struct {
	Dao *system.DictTypeDao
}

func NewDictTypeServiceImpl(dao *system.DictTypeDao) DictTypeService {
	return &DictTypeServiceImpl{
		Dao: dao,
	}
}

// CreateDictType 添加字典类型
func (s *DictTypeServiceImpl) CreateDictType(dto a.AddDictTypeDto) error {
	return s.Dao.CreateDictType(dto)
}

// DeleteDictTypeByIds 删除字典类型
func (s *DictTypeServiceImpl) DeleteDictTypeByIds(ids []int64) error {
	return s.Dao.DeleteDictTypeByIds(ids)
}

// UpdateDictType 更新字典类型
func (s *DictTypeServiceImpl) UpdateDictType(dto a.UpdateDictTypeDto) error {
	return s.Dao.UpdateDictType(dto)
}

// UpdateDictTypeStatus 更新字典类型状态
func (s *DictTypeServiceImpl) UpdateDictTypeStatus(dto a.UpdateDictTypeStatusDto) error {
	return s.Dao.UpdateDictTypeStatus(dto)
}

// QueryDictTypeDetail 查询字典类型详情
func (s *DictTypeServiceImpl) QueryDictTypeDetail(dto a.QueryDictTypeDetailDto) (b.DictType, error) {
	return s.Dao.QueryDictTypeDetail(dto)
}

// QueryDictTypeList 查询字典类型列表
func (s *DictTypeServiceImpl) QueryDictTypeList(dto a.QueryDictTypeListDto) ([]b.DictType, int64) {
	return s.Dao.QueryDictTypeList(dto)
}
