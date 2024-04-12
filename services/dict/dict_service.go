package dict

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/repositories"
)

type DictService struct {
}

func (d *DictService) CreateDict(dto dto.DictDto) error {
	return repositories.CreatDict(dto)
}

func (d *DictService) GetDictList(current int, pageSize int) ([]models.SysDict, int) {
	return repositories.GetDictList(current, pageSize)
}

func (d *DictService) UpdateDict(dictDto dto.DictDto) error {
	return repositories.UpdateDict(dictDto)
}

func (d *DictService) DeleteDictById(id int64) error {
	return repositories.DeleteDictById(id)
}
