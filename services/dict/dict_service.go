package dict

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// DictService 字典操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 11:18
*/
type DictService interface {

	// CreateDict 创建字典
	CreateDict(dto dto.DictDto) error

	// QueryDictList 查询字典列表
	QueryDictList(current int, pageSize int) ([]models.Dict, int64)

	// UpdateDict 更新字典
	UpdateDict(userDto dto.DictDto) error

	// DeleteDictByIds 删除字典
	DeleteDictByIds(ids []int64) error
}
