package dict

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DictContract interface {
	CreateDict(dto dto.DictDto) error

	GetDictList(current int, pageSize int) ([]models.SysDict, int)

	UpdateDict(userDto dto.DictDto) error

	DeleteDictById(id int64) error
}
