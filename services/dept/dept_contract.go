package dept

import (
	"github.com/feihua/simple-go/dto"
)

type DeptContract interface {
	CreateDept(dto dto.DeptDto) error

	UpdateDept(deptDto dto.DeptDto) error

	DeleteDeptById(id int) error
}
