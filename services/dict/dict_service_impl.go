package dict

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DictServiceImpl struct {
}

func (d *DictServiceImpl) CreateDict(dto dto.DictDto) error {
	return dao.CreatDict(dto)
}

func (d *DictServiceImpl) QueryDictList(current int, pageSize int) ([]models.Dict, int) {
	return dao.QueryDictList(current, pageSize)
}

func (d *DictServiceImpl) UpdateDict(dictDto dto.DictDto) error {
	return dao.UpdateDict(dictDto)
}

func (d *DictServiceImpl) DeleteDictByIds(ids []int64) error {
	return dao.DeleteDictByIds(ids)
}
