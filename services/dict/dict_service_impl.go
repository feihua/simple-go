package dict

import (
	"errors"
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// DictServiceImpl 字典操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 11:18
*/
type DictServiceImpl struct {
}

// CreateDict 创建字典
func (d *DictServiceImpl) CreateDict(dto dto.DictDto) error {

	dept, err := dao.QueryDictByName(dto.Value, dto.Type)
	if err != nil {
		return err
	}

	if dept != nil {
		return errors.New("字典已存在")
	}

	return dao.CreateDict(dto)
}

// QueryDictList 查询字典列表
func (d *DictServiceImpl) QueryDictList(current int, pageSize int) ([]models.Dict, int) {
	return dao.QueryDictList(current, pageSize)
}

// UpdateDict 更新字典
func (d *DictServiceImpl) UpdateDict(dictDto dto.DictDto) error {
	return dao.UpdateDict(dictDto)
}

// DeleteDictByIds 删除字典
func (d *DictServiceImpl) DeleteDictByIds(ids []int64) error {
	return dao.DeleteDictByIds(ids)
}
