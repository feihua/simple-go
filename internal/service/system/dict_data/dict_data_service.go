package dict_data

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// DictDataService 字典数据操作接口
type DictDataService interface {
	// CreateDictData 添加字典数据
	CreateDictData(dto d.AddDictDataDto) error
	// DeleteDictDataByIds 删除字典数据
	DeleteDictDataByIds(ids []int64) error
	// UpdateDictData 更新字典数据
	UpdateDictData(dto d.UpdateDictDataDto) error
	// UpdateDictDataStatus 更新字典数据状态
	UpdateDictDataStatus(dto d.UpdateDictDataStatusDto) error
	// QueryDictDataDetail 查询字典数据详情
	QueryDictDataDetail(dto d.QueryDictDataDetailDto) (*d.QueryDictDataListDtoResp, error)
	// QueryDictDataById 根据id查询字典数据详情
	QueryDictDataById(id int64) (*d.QueryDictDataListDtoResp, error)
	// QueryDictDataList 查询字典数据列表
	QueryDictDataList(dto d.QueryDictDataListDto) ([]*d.QueryDictDataListDtoResp, int64, error)
}
