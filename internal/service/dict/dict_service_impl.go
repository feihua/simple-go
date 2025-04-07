package dict

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
)

// DictServiceImpl 字典操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 11:18
*/
type DictServiceImpl struct {
	Dao *dao.DictDao
}

func NewDictServiceImpl(Dao *dao.DictDao) DictService {
	return &DictServiceImpl{Dao: Dao}
}

// CreateDict 创建字典
func (d *DictServiceImpl) CreateDict(dto dto.DictDto) error {

	dept, err := d.Dao.QueryDictByName(dto.Value, dto.Type)
	if err != nil {
		return err
	}

	if dept != nil {
		return errors.New("字典已存在")
	}

	return d.Dao.CreateDict(dto)
}

// QueryDictList 查询字典列表
func (d *DictServiceImpl) QueryDictList(current int, pageSize int) ([]model.Dict, int64) {
	return d.Dao.QueryDictList(current, pageSize)
}

// UpdateDict 更新字典
func (d *DictServiceImpl) UpdateDict(dictDto dto.DictDto) error {
	return d.Dao.UpdateDict(dictDto)
}

// DeleteDictByIds 删除字典
func (d *DictServiceImpl) DeleteDictByIds(ids []int64) error {
	return d.Dao.DeleteDictByIds(ids)
}
