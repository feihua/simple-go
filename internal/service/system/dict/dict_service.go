package dict

import (
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
)

// DictService 字典操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 11:18
*/
type DictService interface {

	// CreateDict 创建字典
	CreateDict(dto system.DictDto) error

	// QueryDictList 查询字典列表
	QueryDictList(current int, pageSize int) ([]system2.Dict, int64)

	// UpdateDict 更新字典
	UpdateDict(userDto system.DictDto) error

	// DeleteDictByIds 删除字典
	DeleteDictByIds(ids []int64) error
}
