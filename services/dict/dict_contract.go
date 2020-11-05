package dict

import (
	"simple-go/dto"
	"simple-go/models"
)

type DictContract interface {

	CreateDict(dto dto.DictDto) error

	GetDictList(current int, pageSize int) ([]models.SysDict,int)

	UpdateDict(userDto dto.DictDto) error

	DeleteDictById(id int64) error
}
