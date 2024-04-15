package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DeptService interface {
	CreateDept(dto dto.DeptDto) error

	QueryDeptList() ([]models.Dept, int)

	UpdateDept(deptDto dto.DeptDto) error

	DeleteDeptByIds(ids []int64) error
}
