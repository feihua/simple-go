package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DeptContract interface {
	CreateDept(dto dto.DeptDto) error

	QueryDeptList() ([]models.Dept, int)

	UpdateDept(deptDto dto.DeptDto) error

	DeleteDeptById(id int64) error
}
