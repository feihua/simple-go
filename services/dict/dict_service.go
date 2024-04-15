package dict

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DictService struct {
}

func (d *DictService) CreateDict(dto dto.DictDto) error {
	return dao.CreatDict(dto)
}

func (d *DictService) QueryDictList(current int, pageSize int) ([]models.Dict, int) {
	return dao.QueryDictList(current, pageSize)
}

func (d *DictService) UpdateDict(dictDto dto.DictDto) error {
	return dao.UpdateDict(dictDto)
}

func (d *DictService) DeleteDictById(id int64) error {
	return dao.DeleteDictById(id)
}
