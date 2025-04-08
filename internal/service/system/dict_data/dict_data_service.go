package dict_data

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// DictDataService 字典数据操作接口
type DictDataService interface {
	// CreateDictData 添加字典数据
	CreateDictData(dto b.AddDictDataDto) error
	// DeleteDictDataByIds 删除字典数据
	DeleteDictDataByIds(ids []int64) error
	// UpdateDictData 更新字典数据
	UpdateDictData(dto b.UpdateDictDataDto) error
	// UpdateDictDataStatus 更新字典数据状态
	UpdateDictDataStatus(dto b.UpdateDictDataStatusDto) error
	// QueryDictDataDetail 查询字典数据详情
	QueryDictDataDetail(dto b.QueryDictDataDetailDto) (a.DictData, error)
	// QueryDictDataList 查询字典数据列表
	QueryDictDataList(dto b.QueryDictDataListDto) ([]a.DictData, int64)
}
