package dict_type

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"time"
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
func (s *DictTypeServiceImpl) CreateDictType(dto d.AddDictTypeDto) error {
	byName, err := s.Dao.QueryDictTypeByName(dto.DictName)
	if err != nil {
		return err
	}
	if byName != nil {
		return errors.New("字典类型名称已存在")
	}

	byType, err := s.Dao.QueryDictTypeByType(dto.DictType)
	if err != nil {
		return err
	}

	if byType != nil {
		return errors.New("字典类型编码已存在")
	}
	return s.Dao.CreateDictType(dto)
}

// DeleteDictTypeByIds 删除字典类型
func (s *DictTypeServiceImpl) DeleteDictTypeByIds(ids []int64) error {
	return s.Dao.DeleteDictTypeByIds(ids)
}

// UpdateDictType 更新字典类型
func (s *DictTypeServiceImpl) UpdateDictType(dto d.UpdateDictTypeDto) error {
	item, err := s.Dao.QueryDictTypeById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("字典类型不存在")
	}

	byName, err := s.Dao.QueryDictTypeByName(dto.DictName)
	if err != nil {
		return err
	}
	if byName != nil && item.Id != dto.Id {
		return errors.New("字典类型名称已存在")
	}

	byType, err := s.Dao.QueryDictTypeByType(dto.DictType)
	if err != nil {
		return err
	}

	if byType != nil && item.Id != dto.Id {
		return errors.New("字典类型编码已存在")
	}

	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateDictType(dto)
}

// UpdateDictTypeStatus 更新字典类型状态
func (s *DictTypeServiceImpl) UpdateDictTypeStatus(dto d.UpdateDictTypeStatusDto) error {
	return s.Dao.UpdateDictTypeStatus(dto)
}

// QueryDictTypeDetail 查询字典类型详情
func (s *DictTypeServiceImpl) QueryDictTypeDetail(dto d.QueryDictTypeDetailDto) (*d.QueryDictTypeListDtoResp, error) {
	item, err := s.Dao.QueryDictTypeDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("字典类型不存在")
	}

	return &d.QueryDictTypeListDtoResp{
		Id:         item.Id,                             // 字典主键
		DictName:   item.DictName,                       // 字典名称
		DictType:   item.DictType,                       // 字典类型
		Status:     item.Status,                         // 状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryDictTypeById 根据id查询字典类型详情
func (s *DictTypeServiceImpl) QueryDictTypeById(id int64) (*d.QueryDictTypeListDtoResp, error) {
	item, err := s.Dao.QueryDictTypeById(id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("字典类型不存在")
	}

	return &d.QueryDictTypeListDtoResp{
		Id:         item.Id,                             // 字典主键
		DictName:   item.DictName,                       // 字典名称
		DictType:   item.DictType,                       // 字典类型
		Status:     item.Status,                         // 状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryDictTypeList 查询字典类型列表
func (s *DictTypeServiceImpl) QueryDictTypeList(dto d.QueryDictTypeListDto) ([]*d.QueryDictTypeListDtoResp, int64, error) {
	result, i, err := s.Dao.QueryDictTypeList(dto)

	if err != nil {
		return nil, 0, err
	}

	list := make([]*d.QueryDictTypeListDtoResp, 0)

	for _, item := range result {
		resp := &d.QueryDictTypeListDtoResp{
			Id:         item.Id,                             // 字典主键
			DictName:   item.DictName,                       // 字典名称
			DictType:   item.DictType,                       // 字典类型
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
