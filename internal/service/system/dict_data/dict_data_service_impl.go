package dict_data

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"time"
)

// DictDataServiceImpl 字典数据操作实现
type DictDataServiceImpl struct {
	Dao         *system.DictDataDao
	DictTypeDao *system.DictTypeDao
}

func NewDictDataServiceImpl(dao *system.DictDataDao, DictTypeDao *system.DictTypeDao) DictDataService {
	return &DictDataServiceImpl{
		Dao:         dao,
		DictTypeDao: DictTypeDao,
	}
}

// CreateDictData 添加字典数据
func (s *DictDataServiceImpl) CreateDictData(dto d.AddDictDataDto) error {
	byType, err := s.DictTypeDao.QueryDictTypeByType(dto.DictType)
	if err != nil {
		return err
	}
	if byType == nil {
		return errors.New("字典类型不存在")
	}

	label, err := s.Dao.QueryDictDataByLabel(dto.DictLabel, dto.DictType)
	if err != nil {
		return err
	}
	if label != nil {
		return errors.New("字典标签已存在")
	}

	byValue, err := s.Dao.QueryDictDataByValue(dto.DictValue, dto.DictType)
	if err != nil {
		return err
	}
	if byValue != nil {
		return errors.New("字典键值已存在")
	}

	if dto.IsDefault == "Y" {
		err = s.Dao.UpdateDictDataDefault(dto.DictType)
		if err != nil {
			return err
		}
	}
	return s.Dao.CreateDictData(dto)
}

// DeleteDictDataByIds 删除字典数据
func (s *DictDataServiceImpl) DeleteDictDataByIds(ids []int64) error {
	return s.Dao.DeleteDictDataByIds(ids)
}

// UpdateDictData 更新字典数据
func (s *DictDataServiceImpl) UpdateDictData(dto d.UpdateDictDataDto) error {
	item, err := s.Dao.QueryDictDataById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("字典数据不存在")
	}

	byType, err := s.DictTypeDao.QueryDictTypeByType(dto.DictType)
	if err != nil {
		return err
	}
	if byType == nil {
		return errors.New("字典类型不存在")
	}

	label, err := s.Dao.QueryDictDataByLabel(dto.DictLabel, dto.DictType)
	if err != nil {
		return err
	}
	if label != nil && item.Id != dto.Id {
		return errors.New("字典标签已存在")
	}

	byValue, err := s.Dao.QueryDictDataByValue(dto.DictValue, dto.DictType)
	if err != nil {
		return err
	}
	if byValue != nil && item.Id != dto.Id {
		return errors.New("字典键值已存在")
	}

	if dto.IsDefault == "Y" {
		err = s.Dao.UpdateDictDataDefault(dto.DictType)
		if err != nil {
			return err
		}
	}
	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateDictData(dto)
}

// UpdateDictDataStatus 更新字典数据状态
func (s *DictDataServiceImpl) UpdateDictDataStatus(dto d.UpdateDictDataStatusDto) error {
	return s.Dao.UpdateDictDataStatus(dto)
}

// QueryDictDataDetail 查询字典数据详情
func (s *DictDataServiceImpl) QueryDictDataDetail(dto d.QueryDictDataDetailDto) (*d.QueryDictDataListDtoResp, error) {
	item, err := s.Dao.QueryDictDataDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("字典数据不存在")
	}

	return &d.QueryDictDataListDtoResp{
		Id:         item.Id,                             // 字典编码
		DictSort:   item.DictSort,                       // 字典排序
		DictLabel:  item.DictLabel,                      // 字典标签
		DictValue:  item.DictValue,                      // 字典键值
		DictType:   item.DictType,                       // 字典类型
		CssClass:   item.CssClass,                       // 样式属性（其他样式扩展）
		ListClass:  item.ListClass,                      // 表格回显样式
		IsDefault:  item.IsDefault,                      // 是否默认（Y是 N否）
		Status:     item.Status,                         // 状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryDictDataById 根据id查询字典数据详情
func (s *DictDataServiceImpl) QueryDictDataById(id int64) (*d.QueryDictDataListDtoResp, error) {
	item, err := s.Dao.QueryDictDataById(id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("字典数据不存在")
	}

	return &d.QueryDictDataListDtoResp{
		Id:         item.Id,                             // 字典编码
		DictSort:   item.DictSort,                       // 字典排序
		DictLabel:  item.DictLabel,                      // 字典标签
		DictValue:  item.DictValue,                      // 字典键值
		DictType:   item.DictType,                       // 字典类型
		CssClass:   item.CssClass,                       // 样式属性（其他样式扩展）
		ListClass:  item.ListClass,                      // 表格回显样式
		IsDefault:  item.IsDefault,                      // 是否默认（Y是 N否）
		Status:     item.Status,                         // 状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryDictDataList 查询字典数据列表
func (s *DictDataServiceImpl) QueryDictDataList(dto d.QueryDictDataListDto) ([]*d.QueryDictDataListDtoResp, int64, error) {
	result, i, err := s.Dao.QueryDictDataList(dto)

	if err != nil {
		return nil, 0, err
	}

	var list []*d.QueryDictDataListDtoResp

	for _, item := range result {
		resp := &d.QueryDictDataListDtoResp{
			Id:         item.Id,                             // 字典编码
			DictSort:   item.DictSort,                       // 字典排序
			DictLabel:  item.DictLabel,                      // 字典标签
			DictValue:  item.DictValue,                      // 字典键值
			DictType:   item.DictType,                       // 字典类型
			CssClass:   item.CssClass,                       // 样式属性（其他样式扩展）
			ListClass:  item.ListClass,                      // 表格回显样式
			IsDefault:  item.IsDefault,                      // 是否默认（Y是 N否）
			Status:     item.Status,                         // 状态（0：停用，1:正常）
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		}

		list = append(list, resp)
	}

	return list, i, nil
}
