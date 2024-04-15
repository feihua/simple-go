package dict

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DictService interface {
	CreateDict(dto dto.DictDto) error

	QueryDictList(current int, pageSize int) ([]models.Dict, int)

	UpdateDict(userDto dto.DictDto) error

	DeleteDictByIds(ids []int64) error
}
