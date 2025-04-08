package dict_data

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// DictDataServiceImpl 字典数据操作实现
type DictDataServiceImpl struct {
	Dao *system.DictDataDao
}

func NewDictDataServiceImpl(dao *system.DictDataDao) DictDataService {
	return &DictDataServiceImpl{
		Dao: dao,
	}
}

// CreateDictData 添加字典数据
func (s *DictDataServiceImpl) CreateDictData(dto a.AddDictDataDto) error {
	return s.Dao.CreateDictData(dto)
}

// DeleteDictDataByIds 删除字典数据
func (s *DictDataServiceImpl) DeleteDictDataByIds(ids []int64) error {
	return s.Dao.DeleteDictDataByIds(ids)
}

// UpdateDictData 更新字典数据
func (s *DictDataServiceImpl) UpdateDictData(dto a.UpdateDictDataDto) error {
	return s.Dao.UpdateDictData(dto)
}

// UpdateDictDataStatus 更新字典数据状态
func (s *DictDataServiceImpl) UpdateDictDataStatus(dto a.UpdateDictDataStatusDto) error {
	return s.Dao.UpdateDictDataStatus(dto)
}

// QueryDictDataDetail 查询字典数据详情
func (s *DictDataServiceImpl) QueryDictDataDetail(dto a.QueryDictDataDetailDto) (b.DictData, error) {
	return s.Dao.QueryDictDataDetail(dto)
}

// QueryDictDataList 查询字典数据列表
func (s *DictDataServiceImpl) QueryDictDataList(dto a.QueryDictDataListDto) ([]b.DictData, int64) {
	return s.Dao.QueryDictDataList(dto)
}
