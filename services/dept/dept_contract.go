package dept

import (
	"simple-go/dto"
	"simple-go/pkg/util"
)

type DeptContract interface {

	CreateDept(dto dto.DeptDto) error

	GetDeptList() []util.Tree

	UpdateDept(deptDto dto.DeptDto) error

	DeleteDeptById(id int) error
}
